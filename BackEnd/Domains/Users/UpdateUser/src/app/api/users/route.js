import { NextResponse } from "next/server";
import { conn } from "@/libs/mysql";


// Update a user
export async function PUT(request) {
  try {
    const { id, name, email, password } = await request.json();
    await conn.query("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", [name, email, password, id]);
    return NextResponse.json({ message: "Usuario actualizado exitosamente" });
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
