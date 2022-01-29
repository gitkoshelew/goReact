import React from 'react';
import { Story, Meta } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { ButtonProps, Button } from './Button';

const meta: Meta = {
    title: 'Button',
    component: Button,
}
export default meta

export const Template: Story<ButtonProps> = (args) => <Button {...args} />;

export const OrderButton = Template.bind({});
export const UploadButton = Template.bind({});

OrderButton.args = {
    view: 'order',
    onClick: action('OrderButton click')
}
UploadButton.args = {
    view: 'upload',
    disabled: true
}



