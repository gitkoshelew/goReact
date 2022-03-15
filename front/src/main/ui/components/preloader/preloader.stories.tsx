import React from 'react';
import { Story, Meta } from '@storybook/react';
import Preloader from "./preloader";


const meta: Meta = {
    title: 'Preloader',
    component: Preloader,
}
export default meta

export const Template: Story = (args) => <Preloader {...args}   />

