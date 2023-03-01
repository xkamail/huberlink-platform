import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'

const SignUpPage = () => {
  return (
    <div className="container mx-auto py-10 h-screen bg-slate-100 shadow">
      <div className="rounded-lg text-center bg-white p-4 max-w-2xl mx-auto">
        <h1 className="mb-10">
          <span className="text-2xl font-bold ">
            Create account by using discord
          </span>
        </h1>
        <Button
          variant="discord"
          to="https://discord.com/api/oauth2/authorize?client_id=1079964496866582578&redirect_uri=https%3A%2F%2Fhuberlink.vercel.app%2Fcallback&response_type=code&scope=identify%20email"
        >
          <Icons.discord className="w-5 h-5 text-slate-100 mr-2" />
          Continute with discord
        </Button>
      </div>
    </div>
  )
}

export default SignUpPage

SignUpPage.displayName = 'SignUpPage'
