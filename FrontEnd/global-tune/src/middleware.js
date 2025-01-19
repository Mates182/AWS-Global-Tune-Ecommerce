import { NextResponse } from "next/server";
import { jwtVerify } from "jose";
export async function middleware(request) {
  const pathname = request.nextUrl.pathname;
  const token = request.cookies.get("token");
  if (pathname.includes("/admin")) {
    if (token === undefined) {
      return NextResponse.redirect(new URL("/login", request.url));
    }
    try {
      const { payload } = await jwtVerify(token.value, new TextEncoder().encode( process.env.SECRET_KEY));
      if(payload.role==="admin"){
        return NextResponse.next();
      
      }
      return NextResponse.redirect(new URL("/login", request.url));
    } catch (error) {
      console.log(error)
      return NextResponse.redirect(new URL("/login", request.url));
    }
  }
  return NextResponse.next();
}
