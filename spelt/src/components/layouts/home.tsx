import { cx } from 'class-variance-authority'
import { Button, buttonVariants } from '@/components/ui/button'
import { useTheme } from '@/context/theme-provider'

const Home = () => {
  const { theme, setTheme } = useTheme()
  return (
    <div className="flex min-h-svh flex-col items-center justify-center">
      <Button
        className={cx(buttonVariants({ variant: 'default', size: 'sm' }))}
        onClick={() => {
          console.log('toggle theme', theme)
          if (theme === 'system') {
            setTheme('dark')
            return
          }
          setTheme(theme === 'light' ? 'dark' : 'light')
          console.log('new theme', theme)
        }}
      >
        Click me
      </Button>
    </div>
  )
}

export { Home }
