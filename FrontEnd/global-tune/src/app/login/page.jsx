"use client";
import Image from "next/image";
import React, { use, useState } from "react";
import { loginAuth, logout } from "@/services/Rest/login";
import { BsGlobeAmericas } from "react-icons/bs";

function LoginPage() {
  const [isAuth, setIsAuth] = useState(false);
  const handleSubmit = async (formData) => {
    const credentials = {
      email: formData.get("email"),
      password: formData.get("password"),
    };
    try {
      const response = await loginAuth(credentials);
      if (response.message == "Login successful") {
        setIsAuth(true);
      } else {
        setIsAuth(false);
      }
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  };
  const handleLogout = async () => {
    try {
      const response = await logout();
      if (response.message == "Logged out successfully") {
        setIsAuth(false);
      }
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <div>
      {isAuth ? (
        <div>
          <h1>Logged in</h1>
          <button onClick={handleLogout}>Logout</button>
        </div>
      ) : (
        <section className="vh-100">
          <div className="container py-5 h-100">
            <div className="row d-flex justify-content-center align-items-center h-100">
              <div className="col col-xl-10">
                <div className="card" style={{ borderRadius: "1rem" }}>
                  <div className="row g-0">
                    <div className="col-md-6 col-lg-5 d-none d-md-block">
                      <Image
                        src="/login-image.png"
                        alt="Login Image"
                        width={500}
                        height={500}
                        className="img-fluid"
                        style={{ borderRadius: "1rem 0 0 1rem" }}
                      />
                    </div>
                    <div className="col-md-6 col-lg-7 d-flex align-items-center">
                      <div className="card-body p-4 p-lg-5 text-black">
                        <form action={handleSubmit}>
                          <div className="d-flex align-items-center mb-3 pb-1">
                            <i
                              className="fas fa-cubes fa-2x me-3"
                              style={{ color: "#ff6219", fontSize: "38px" }}
                            >
                              <BsGlobeAmericas />
                            </i>
                            <span className="h1 fw-bold mb-0">GlobalTune</span>
                          </div>
                          <h5
                            className="fw-normal mb-3 pb-3"
                            style={{ letterSpacing: "1px" }}
                          >
                            Sign into your account
                          </h5>
                          <div className="form-outline mb-4">
                            <input
                              name="email"
                              type="email"
                              className="form-control form-control-lg"
                              required
                            />
                            <label className="form-label">Email address</label>
                          </div>
                          <div className="form-outline mb-4">
                            <input
                              name="password"
                              type="password"
                              className="form-control form-control-lg"
                              required
                            />
                            <label className="form-label">Password</label>
                          </div>
                          <div className="pt-1 mb-4">
                            <button
                              className="btn btn-dark btn-lg btn-block"
                            >
                              Login
                            </button>
                          </div>
                          <a className="small text-muted" href="#!">
                            Forgot password?
                          </a>
                          <p
                            className="mb-5 pb-lg-2"
                            style={{ color: "#393f81" }}
                          >
                            Don't have an account?{" "}
                            <a href="#!" style={{ color: "#393f81" }}>
                              Register here
                            </a>
                          </p>

                          <a href="#!" className="small text-muted">
                            Terms of use.
                          </a>
                          <a href="#!" className="small text-muted">
                            Privacy policy
                          </a>
                        </form>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      )}
    </div>
  );
}

export default LoginPage;
