import React from 'react';
import { Story, Meta } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { BtnPropsType, Button } from './Button';

const meta: Meta = {
    title: 'Button',
    component: Button,
}
export default meta

export const Template: Story<BtnPropsType> = (args) => <Button {...args} />;

export const BuyButton = Template.bind({});
export const OrderButton = Template.bind({});
export const UploadButton = Template.bind({});

BuyButton.args = {
    type: 'buy',
    isActive: true
}

OrderButton.args = {
    type: 'order',
    isActive: true,
    onClick: action('OrderButton click')
}
UploadButton.args = {
    type: 'upload',
    isActive: true
}



