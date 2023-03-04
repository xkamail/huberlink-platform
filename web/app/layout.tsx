import { Toaster } from '@/components/ui/toaster'
import { UserContextProvider } from '@/lib/contexts/UserContext'
import { cn } from '@/lib/utils'
import '@/styles/output.css'
import { Sarabun } from '@next/font/google'

const inter = Sarabun({
  weight: ['100', '200', '300', '400', '500', '600', '700', '800'],
  subsets: ['latin-ext'],
})

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
        <div className="flex min-h-screen flex-col">
          <UserContextProvider>{children}</UserContextProvider>
        </div>
        <Toaster />
      </body>
    </html>
  )
}
