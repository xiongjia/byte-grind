import { createFileRoute } from '@tanstack/react-router'
import { Home } from '@/components/layouts/home'

export const Route = createFileRoute('/')({
  component: Home,
})
