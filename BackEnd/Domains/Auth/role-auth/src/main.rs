use jsonwebtoken::{decode, DecodingKey, Validation, errors::ErrorKind};
use serde::{Deserialize, Serialize};
use tonic::{transport::Server, Request, Response, Status};

pub mod auth {
    tonic::include_proto!("auth");
}

use auth::{
    auth_service_server::{AuthService, AuthServiceServer},
    ValidateTokenRequest, ValidateTokenResponse,
};

const SECRET_KEY: &str = "my_secret_key"; // Change in production

#[derive(Debug, Serialize, Deserialize)]
struct Claims {
    sub: String,      
    role: String,   
    user_id: String,  
    exp: usize,      
}

#[derive(Default)]
pub struct AuthServiceImpl;

#[tonic::async_trait]
impl AuthService for AuthServiceImpl {
    async fn validate_token(
        &self,
        request: Request<ValidateTokenRequest>,
    ) -> Result<Response<ValidateTokenResponse>, Status> {
        let req = request.into_inner();

        let mut validation = Validation::default();
        validation.validate_exp = true;

        let token_data = decode::<Claims>(
            &req.token,
            &DecodingKey::from_secret(SECRET_KEY.as_ref()),
            &validation,
        );

        match token_data {
            Ok(data) => {
                if data.claims.role == req.required_role {
                    Ok(Response::new(ValidateTokenResponse {
                        valid: true,
                        user_id: data.claims.sub,
                        message: "Valid Token".into(),
                    }))
                } else {
                    Ok(Response::new(ValidateTokenResponse {
                        valid: false,
                        user_id: "".into(),
                        message: "Invalid Role".into(),
                    }))
                }
            }
            Err(e) => {
                println!("Error decoding token: {:?}", e);
                if let ErrorKind::ExpiredSignature = e.kind() {
                    return Ok(Response::new(ValidateTokenResponse {
                        valid: false,
                        user_id: "".into(),
                        message: "Token Expired".into(),
                    }));
                }
                Ok(Response::new(ValidateTokenResponse {
                    valid: false,
                    user_id: "".into(),
                    message: "Invalid Token".into(),
                }))
            }
        }
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let auth_service = AuthServiceImpl::default();

    println!("AuthService runnin on port {}", addr);

    Server::builder()
        .add_service(AuthServiceServer::new(auth_service))
        .serve(addr)
        .await?;

    Ok(())
}
