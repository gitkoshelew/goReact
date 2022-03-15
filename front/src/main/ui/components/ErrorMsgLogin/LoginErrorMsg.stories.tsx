import React from 'react';
import { Story, Meta } from '@storybook/react';
import { LoginErrorMsg, LoginErrorMsgPropsType } from "./LoginErrorMsg";


const meta: Meta = {
  title: 'LoginErrorMsg',
  component: LoginErrorMsg,
}
export default meta

export const Template: Story<LoginErrorMsgPropsType> = (args) => <LoginErrorMsg {...args}   />
export const BasicError = Template.bind({});

BasicError.args ={
  ErrorMsg: 'ERROR!!!'
}
