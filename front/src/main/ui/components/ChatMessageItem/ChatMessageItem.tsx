import React from 'react'
import s from './ChatMessageItem.module.scss'

const { messageSender, messageText } = s

type PropsType = {
  sender: number
  text: string
}

export const ChatMessageItem = ({ sender, text }: PropsType) => {
  return (
    <div>
      <div className={messageSender}>{sender}</div>
      <div className={messageText}>{text}</div>
    </div>
  )
}
