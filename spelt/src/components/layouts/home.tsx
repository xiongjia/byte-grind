import { Main } from '@/components/layouts/main'
import { Overview } from '@/components/layouts/overview'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useTheme } from '@/context/theme-provider'
import { useTranslation } from 'react-i18next'

const Home = () => {
  const { theme, setTheme } = useTheme()
  const { t, i18n } = useTranslation()

  return (
    <Main>
      <div className="grid grid-cols-1 gap-4 lg:grid-cols-7">
        <Card className="col-span-1 lg:col-span-4">
          <CardHeader>
            <CardTitle>Overview</CardTitle>
          </CardHeader>
          <CardContent className="ps-2">
            <Overview />
          </CardContent>
        </Card>
      </div>

      <div className="flex min-h-svh flex-col items-center justify-center">
        <h1>{t('Welcome to React')}</h1>
        <Button
          onClick={() => {
            console.log('lang ', i18n.language)
            if (i18n.language === 'cn') {
              i18n.changeLanguage('en')
            } else {
              i18n.changeLanguage('cn')
            }
          }}
        >
          Change Lang
        </Button>
        <br />
        <Button
          variant={'default'}
          size={'sm'}
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
    </Main>
  )
}

export { Home }
