import type { Meta, StoryObj } from '@storybook/react-vite'
import { Button } from '@/components/ui/button'

const meta = {
  title: 'general/button',
  component: Button,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Button>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: { children: 'Click me', variant: 'default', size: 'default' },
  decorators: [
    (Story) => (
      <>
        <Story />
      </>
    ),
  ],
}

export const Secondary: Story = {
  args: { children: 'Click me', variant: 'secondary', size: 'default' },
}
