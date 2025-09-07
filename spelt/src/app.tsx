import { cx } from 'class-variance-authority'
import { Button, buttonVariants } from '@/components/button'

function App() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center">
      <Button
        className={cx(
          buttonVariants({ variant: 'destructive', size: 'sm' }),
          'hover:underline',
        )}
      >
        Click me
      </Button>
    </div>
  )
}

export default App
