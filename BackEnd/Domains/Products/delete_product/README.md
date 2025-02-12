# DeleteProduct

A Gin-based microservice created with ginshot.

## Structure

```bash
delete_product/
├── app/
│   └── main.go
├── config/
│   ├── cors/
│   │   └── cors.go
│   └── kafka/
│       └── offset_manager.go
├── deployments/
│   └── docker-compose.yml
├── docs/
├── internal/
│   ├── auth/
│   │   ├── auth.pb.go
│   │   ├── auth_client.go
│   │   ├── auth_grpc.pb.go
│   │   └── validate_token.go
│   ├── bootstrap/
│   │   └── setup/
│   │       └── setup.go
│   ├── controller/
│   │   └── delete_product-controller.go
│   ├── data/
│   │   ├── messages/
│   │   │   ├── product_topic_message.go
│   │   │   └── replication_topic_message.go
│   │   ├── models/
│   │   │   └── product.go
│   │   ├── requests/
│   │   │   └── delete_product_request.go
│   │   └── responses/
│   │       └── delete_product_response.go
│   ├── event/
│   │   ├── consumer.go
│   │   ├── producer.go
│   │   └── start_consumer.go
│   ├── middleware/
│   ├── repository/
│   │   ├── dbcontext.go
│   │   ├── product_replicator.go
│   │   └── product_repository.go
│   ├── secrets/
│   │   └── products-env.go
│   └── service/
│       ├── delete_product-service-impl.go
│       └── delete_product-service.go
├── pkg/
├── router/
│   └── router.go
├── test/
├── .dockerignore
├── .env
├── Dockerfile
├── README.md
├── go.mod
└── go.sum
```