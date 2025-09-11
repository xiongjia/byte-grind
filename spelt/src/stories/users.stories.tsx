import type { Meta, StoryObj } from '@storybook/react-vite'
import { Users } from '@/components/layouts/users/users'

const meta = {
  title: 'app/users',
  component: Users,
  parameters: {
    layout: 'centered',
  },
} satisfies Meta<typeof Users>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {},
}
