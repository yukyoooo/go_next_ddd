import { ComponentMeta, ComponentStory } from '@storybook/react'
import { ThemeProvider } from 'styled-components'
import { theme } from '../../../themes'

import Button from './index'

export default {
  title: 'atoms/Button',
  argTypes: {
    variant: {
      options: ['primary', 'secondary'],
      control: { type: 'radio' },
      defaultValue: 'primary',
      description: 'ボタンバリアント',
      table: {
        type: { summary: 'primary | secondary' },
        defaultValue: { summary: 'primary' },
      },
    },
    children: {
      control: { type: 'text' },
      defaultValue: 'ボタン',
      description: 'ボタンのラベル',
      table: {
        type: { summary: 'string' },
      },
    },
    disabled: {
      control: { type: 'boolean' },
      defaultValue: false,
      description: 'ボタンを無効化する',
      table: {
        type: { summary: 'boolean' },
      },
    },
    width: {
      control: { type: 'number' },
      description: 'ボタンの横幅',
      table: {
        type: { summary: 'number' },
      },
    },
    height: {
      control: { type: 'number' },
      description: 'ボタンの縦幅',
      table: {
        type: { summary: 'number' },
      },
    },
    onClick: {
      description: 'ボタンをクリックしたときのイベント',
      table: {
        type: { summary: '() => void' },
      },
    },
  },
} as ComponentMeta<typeof Button>

const Template: ComponentStory<typeof Button> = (args) => {
  return (
    <ThemeProvider theme={theme}>
      <Button {...args} />
    </ThemeProvider>
  )
}

export const Primary = Template.bind({})
Primary.args = { variant: 'primary', children: 'Primary button' }

export const Secondary = Template.bind({})
Secondary.args = { variant: 'secondary', children: 'Secondary button' }

export const Disabled = Template.bind({})
Disabled.args = { disabled: true, children: 'Disabled button' }
