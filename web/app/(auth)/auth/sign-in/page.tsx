import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Separator } from '@/components/ui/sperator'

const SignUpPage = () => {
  return (
    <div className=" w-full py-10 h-screen bg-slate-100 shadow">
      <div className="rounded-lg text-center bg-white p-4 max-w-2xl mx-auto">
        <div className="w-80 mx-auto">
          <h1 className="mb-6">
            <span className="text-2xl font-bold ">Sign in</span>
          </h1>
          <Form className=" mb-5 mx-auto">
            <Input placeholder="Username or E-mail" type="text" />
            <Input placeholder="Password" type="password" />
            <Button block type="submit">
              Sign in
            </Button>
          </Form>
          <Separator />
          <Button
            className="mt-4"
            variant="discord"
            to="https://discord.com/api/oauth2/authorize?client_id=1079964496866582578&redirect_uri=https%3A%2F%2Fhuberlink.vercel.app%2Fcallback&response_type=code&scope=identify%20email"
            block
          >
            <Icons.discord className="w-5 h-5 text-slate-100 mr-2" />
            Continute with Discord
          </Button>
        </div>
      </div>
    </div>
  )
}

export default SignUpPage

SignUpPage.displayName = 'SignUpPage'
