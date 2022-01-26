import React from 'react';
import { Story, Meta } from '@storybook/react';
import { SelectUI } from "./Select";


const meta: Meta = {
    title: 'SelectUI',
    component: SelectUI,
}
export default meta

export const Template: Story = (args) => <SelectUI {...args}/>
