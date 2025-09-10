import type { Preview, ReactRenderer } from '@storybook/react-vite'
import '../src/index.css'
import { withThemeByClassName } from '@storybook/addon-themes'

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
    a11y: {
      // 'todo' - show a11y violations in the test UI only
      // 'error' - fail CI on a11y violations
      // 'off' - skip a11y checks entirely
      test: 'todo',
    },
    decorators: [
      withThemeByClassName<ReactRenderer>({
        themes: {
          light: 'light',
          dark: 'dark',
        },
        defaultTheme: 'light',
      }),
    ],
  },
}

export default preview
