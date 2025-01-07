"use client";
import React from "react";
import {loginAuth} from '@/services/Rest/login'

function LoginPage() {
  const handleSubmit = async (formData) => {
    const credentials = {
      email: formData.get("email"),
      password: formData.get("password"),
    };
    try {
      const response = await loginAuth(credentials);
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <div>
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
    </div>
  );
}

export default LoginPage;
