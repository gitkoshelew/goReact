import React, { useMemo } from 'react'
import s from './ChatMessageItem.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'

const { messageItem, myMessageItem, messageSender, messageText } = s

type PropsType = {
  sender: number
  text: string
}

export const ChatMessageItem = ({ sender, text }: PropsType) => {
  const me = useSelector((state: AppRootState) => state.LoginPage.user)

  const myMessage = useMemo(() => {
    if (me) {
      return sender === me.userId
    }
  }, [me])

  return (
    <div className={`${messageItem} ${myMessage ? myMessageItem : ''}`}>
      <h3 className={messageSender}>{sender}</h3>
      <div className={messageText}>{text}</div>
    </div>
  )
}
