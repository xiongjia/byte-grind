import { ThemeProvider } from '@/components/theme-provider'
import { Home } from './home'

function App() {
  return (
    <ThemeProvider defaultTheme="light">
      <Home />
    </ThemeProvider>
  )
}

export default App
