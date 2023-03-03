import { NextResponse } from 'next/server'
// middleware.ts
import type { NextRequest } from 'next/server'
import { fetchy } from './services/rawFetch'

// This function can be marked `async` if using `await` inside
export async function middleware(req: NextRequest) {
  const baseURL = req.url
  const accessToken = req.cookies.get('accessToken')
  if (!accessToken) {
    return NextResponse.redirect(
      new URL('/auth/sign-in?redirect=' + req.nextUrl.pathname, baseURL)
    )
  }
  const res = await fetchy.get(req, `/auth/me`) // <--- This is the problem
  if (!res.success) {
    return NextResponse.redirect(
      new URL('/auth/sign-in?redirect=' + req.nextUrl.pathname, baseURL)
    )
  }
  const r = NextResponse.next()
  const _accessToken = req.cookies.get('accessToken')?.value
  if (_accessToken) {
    r.cookies.set('accessToken', _accessToken)
  }

  const _refreshToken = req.cookies.get('refreshToken')?.value
  if (_refreshToken) {
    r.cookies.set('refreshToken', _refreshToken)
  }
  return r
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: ['/h/:path*', '/account/:path*'],
}
