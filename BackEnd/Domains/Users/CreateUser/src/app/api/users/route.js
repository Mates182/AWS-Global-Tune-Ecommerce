import { NextResponse } from "next/server";
import { conn } from "@/libs/mysql";

// Create a new user
export async function POST(request) {
  try {
    const { name, email, password, role } = await request.json();
    const result = await conn.query("INSERT INTO users SET ?", {
      name,
      email,
      password,
      role,
    });

    return NextResponse.json({
      name,
      email,
      role,
      id: result.insertId,
    });
  } catch (error) {
    console.log(error);
    return NextResponse.json(
      {
        message: error.message,
      },
      {
        status: 500,
      }
    );
  }
}
