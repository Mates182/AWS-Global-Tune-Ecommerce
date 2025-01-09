"use client";
import React, { use, useState } from "react";
import { loginAuth, logout } from "@/services/Rest/login";

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
        <form action={handleSubmit}>
          <input name="email" type="email" placeholder="email" required />
          <input
            name="password"
            type="password"
            placeholder="password"
            required
          />
          <button>login</button>
        </form>
      )}
    </div>
  );
}

export default LoginPage;
