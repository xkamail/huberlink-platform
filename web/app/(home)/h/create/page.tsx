import TopNavigation from '@/components/navigation/TopNavigation'
import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {useForm} from "react-hook-form";

const HomePage = () => {

  const ctx = useForm()

  const submit = (data: ICreateHomeForm) => {
    console.log(data)
  }

  return (
    <div>
      <TopNavigation title="Create home" />
      <div className=" max-w-xl mx-auto">
        <div className="bg-white rounded-lg p-4">
          <div>
            <Form ctx={ctx} onSubmit={submit}>
              <p className="text-sm text-slate-500 dark:text-slate-400">
                Once upon a time, in a far-off land, there was a very lazy king
                who spent all day lounging on his throne. One day, his advisors
                came to him with a problem: the kingdom was running out of
                money.
              </p>
              <Input placeholder="Enter home name" />
              <div className="flex justify-between items-center">
                <Button variant="destructive" type="reset" to="/h">
                  Cancel
                </Button>
                <Button variant="default" type="submit">
                  Create
                </Button>
              </div>
            </Form>
          </div>
        </div>
      </div>
    </div>
  )
}
export default HomePage
