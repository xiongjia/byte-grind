import { Button } from '@/components/button'

function App() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center">
      <Button variant="outline" size="lg" onClick={() => {console.log('click')}}>Click me</Button>
    </div>
  )
}

export default App
