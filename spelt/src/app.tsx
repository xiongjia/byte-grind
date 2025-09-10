import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { createRouter, RouterProvider } from '@tanstack/react-router'
import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import { ThemeProvider } from '@/context/theme-provider'
import { routeTree } from './routeTree.gen'

const resources = {
  en: {
    translation: {
      'Welcome to React': 'Welcome to React',
    },
  },
  cn: {
    translation: {
      'Welcome to React': '你好 React',
    },
  },
}

i18n
  .use(initReactI18next)
  .init({ resources, lng: 'en', interpolation: { escapeValue: false } })

const router = createRouter({ routeTree })
const queryClient = new QueryClient()

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider defaultTheme="light">
        <RouterProvider router={router} />
      </ThemeProvider>
    </QueryClientProvider>
  )
}

export default App
