import { NextResponse } from "next/server";
import { conn } from "@/libs/mysql";

// Get all users
export async function GET() {
  try {
    const results = await conn.query("SELECT * FROM users");
    return NextResponse.json(results);
  } catch (error) {
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