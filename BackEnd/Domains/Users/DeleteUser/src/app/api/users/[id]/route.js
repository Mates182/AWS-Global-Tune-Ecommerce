import { NextResponse } from 'next/server';
import { conn } from '@/libs/mysql';

// Delete a user by ID
export async function DELETE(request, { params }) {
  try {
    await conn.query("DELETE FROM users WHERE id = ?", [params.id]);
    return NextResponse.json({ message: "Usuario eliminado exitosamente" });
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


