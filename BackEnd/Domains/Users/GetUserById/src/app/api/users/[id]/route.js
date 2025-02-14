import { NextResponse } from 'next/server';
import { conn } from '@/libs/mysql';

// Get a user by ID
export async function GET(request, { params }) {
  try {
    const result = await conn.query('SELECT * FROM users WHERE id = ?', [params.id]);

    if (result.length === 0) {
      return NextResponse.json(
        {
          message: "User Not Found",
        },
        {
          status: 404,
        }
      );
    }

    return NextResponse.json(result[0]);
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
