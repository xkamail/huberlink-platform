import LandingNavbar from '@/components/landing/navbar'
import { Button } from '@/components/ui/button'
import { Network } from 'lucide-react'
import TokenTest from './token'

export default function Home() {
  return (
    <>
      <LandingNavbar />
      <div className="container mx-auto">
        <section className="py-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="rounded p-8">
              <h1 className="text-3xl font-extrabold leading-tight tracking-tighter md:text-5xl lg:text-6xl lg:leading-[1.1]">
                Huber<span className="text-indigo-500">Link</span> IoT Solution
                Platform
              </h1>
              <p className="text-slate-700">
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Quisquam
              </p>
              <div className="mt-10 gap-2 flex flex-col lg:flex-row">
                <Button to="/auth/sign-up" variant="primary" size="lg">
                  <span className="text-sm uppercase">Try now</span>
                </Button>
                <Button to="/documentation" variant="subtle" size="lg">
                  <span className="text-sm uppercase">Documentation</span>
                </Button>
                <TokenTest />
              </div>
            </div>
            <div className="flex flex-col p-8 justify-between items-center">
              <div className="w-[234px] rounded-[2rem] h-[507px] bg-slate-100 bg-gradient-to-t border-[10px] border-slate-800 flex items-center justify-center flex-col relative">
                <span className="border border-slate-900 bg-slate-900 w-16 h-4 mt-2 rounded-full absolute top-0"></span>

                <div>
                  <Network className="h-10 w-10 text-slate-600" />
                </div>
                <div className="text-slate-500">HuberLink</div>
              </div>
            </div>
          </div>
        </section>
      </div>
      <section className="w-full bg-slate-100 p-8">
        <div className="container mx-auto">
          <h1 className="text-center text-2xl font-bold ">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam
          </h1>
        </div>
      </section>
    </>
  )
}
