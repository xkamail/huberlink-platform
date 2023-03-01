import { cn } from '@/lib/utils'
import '@/styles/output.css'
import { Inter } from '@next/font/google'

const inter = Inter({ subsets: ['latin'] })

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      {/*
        <head /> will contain the components returned by the nearest parent
        head.tsx. Find out more at https://beta.nextjs.org/docs/api-reference/file-conventions/head
      */}
      <head>
        <meta charSet="utf-8" />
      </head>
      <body
        className={cn`min-h-screen text-slate-900 antialiased dark:bg-slate-900 dark:text-slate-50 ${inter.className}`}
      >
        <div className="flex min-h-screen flex-col">{children}</div>
      </body>
    </html>
  )
}
