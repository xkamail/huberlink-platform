import { Button } from '@/components/ui/button'

export default function Home() {
  return (
    <section className="py-4">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="rounded p-8">
          <h1 className="text-3xl font-extrabold leading-tight tracking-tighter md:text-5xl lg:text-6xl lg:leading-[1.1]">
            Huber<span className="text-indigo-500">Link</span> IoT Solution
            Platform
          </h1>
          <p className="text-slate-700">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam
          </p>
          <div className="mt-10">
            <Button variant="primary">
              <span className="text-sm uppercase">Try now</span>
            </Button>
          </div>
        </div>
        <div className="flex flex-col justify-between">asd</div>
      </div>
    </section>
  )
}
