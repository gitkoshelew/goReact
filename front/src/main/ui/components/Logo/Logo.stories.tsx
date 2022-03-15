import React from 'react';
import { Story, Meta } from '@storybook/react';
import { Logo } from "./Logo";


const meta: Meta = {
    title: 'Logo',
    component: Logo,
    parameters: {
        layout: 'centered',
    },
}
export default meta

export const Template: Story = (args) => <Logo {...args}   />

