import React from 'react';
import { Story, Meta } from '@storybook/react';
import { BasketButton } from "./BasketBtn";


const meta: Meta = {
    title: 'BasketButton',
    component: BasketButton,
}
export default meta

export const Template: Story = (args) => <BasketButton {...args} />;


