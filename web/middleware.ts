// middleware.ts
import type { NextRequest } from 'next/server'
import { NextResponse } from 'next/server'

// This function can be marked `async` if using `await` inside
export async function middleware(req: NextRequest) {
  if (req.nextUrl.pathname === '/h')
    return NextResponse.redirect(new URL('/h/create', req.url))

  const baseURL = req.url
  const accessToken = req.cookies.get('accessToken')
  if (!accessToken) {
    return NextResponse.redirect(
      new URL('/auth/sign-in?redirect=' + req.nextUrl.pathname, baseURL)
    )
  }

  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/me`, {
    headers: {
      Authorization: `Bearer ${accessToken.value}`,
    },
  }).then((r) => r.json()) // <--- This is the problem
  if (!res.success) {
    const refreshToken = req.cookies.get('refreshToken')?.value
    if (res.code === 5 && refreshToken) {
      console.log('refreshing token')
      const res2 = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/refresh-token?refreshToken=${refreshToken}`
      ).then((r) => r.json())
      if (res2.success) {
        req.cookies.set('accessToken', res2.token)
        req.cookies.set('refreshToken', res2.refreshToken)
      }
    }

    return NextResponse.redirect(
      new URL('/auth/sign-in?redirect=' + req.nextUrl.pathname, baseURL)
    )
  }
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: ['/h/:path*', '/account/:path*'],
}
