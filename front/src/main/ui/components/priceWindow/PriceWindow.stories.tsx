import React from 'react';
import { Story, Meta } from '@storybook/react';
import { PriceWindow, PriceWindowPropsType } from "./PriceWindow";


const meta: Meta = {
    title: 'PriceWindow',
    component: PriceWindow,
    argTypes: {
        price: {
            control: { type: 'range', min: 0, max: 2000, step: 1 },
        },
    },
}
export default meta

export const Template: Story<PriceWindowPropsType> = (args) => <PriceWindow {...args}   />

