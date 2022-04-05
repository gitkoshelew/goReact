import React from 'react';
import { Story, Meta } from '@storybook/react';

import { ErrorMsg } from "./ErrorMsg";

const meta: Meta = {
    title: 'ErrorMsg',
    component: ErrorMsg,
}
export default meta

export const Template: Story = (args) => <ErrorMsg {...args} />;

export const Large = Template.bind({});
Large.parameters = {
    controls: { hideNoControlsWarning: true },
};
