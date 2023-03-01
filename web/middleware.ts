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
    const refreshToken = `6ZVx4qJ2aMgkXhYoq_wvQug1dLM8pithHIzkRw30Mjm-9jZNDgOJPfzteCpgn2pD4ZLaVvGp81CdLYbYLLK5nuIjf-aLecynd9qZ1aLAQMGyc1HyeSsjZV6Uxr2OHkxSdlfwlDR3djKfj1sdMOmCNUP2hVxYx-jqpxBkn8r-kVy3qjtGn6wTGdiKSu6bNKltct9soUcNge_P4RZniGbNGoHQM_Ct1-nFxhaalyIjO3_2JwUmJq7oRXuHY959oROs6sFsycyAnqknicIpnznkHHuZ95z3DIlygkU5FPShDgrA0AhPe6zw9BqVj642jqevLOycCIzLB2Ya2DcgpHcT0ZBCR2svRt2Y31KA2qZ9jBj9haoVC8VcBmcFhwVXWcquJCUtsLVrOQGjpEOZ`
    if (res.code === 5 && refreshToken) {
      console.log('refreshing token')
      const res2 = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/refresh-token?refreshToken=${refreshToken}`,
        {
          method: 'POST',
        }
      ).then((r) => r.json())
      if (res2.success) {
        console.log('refresh token success', res2)

        req.cookies.set('accessToken', res2.token)
        req.cookies.set('refreshToken', res2.refreshToken)
        return
      } else {
        console.log('refresh token failed', res2, refreshToken)
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
