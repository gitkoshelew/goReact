import React, { useEffect } from 'react'
import s from './Conversation.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { ChatMessageItem } from '../ChatMessageItem/ChatMessageItem'
import { useParams } from 'react-router-dom'
import { getConversationRequest } from '../../../bll/reducers/ChatPageReducer/chatPage-saga'

const { conversation, conversationHistory, newMessageForm, newMessageField, sendMessageButton } = s

export const Conversation = () => {
  const dispatch = useAppDispatch()
  const messages = useSelector((state: AppRootState) => state.ChatPage.messages)
  const me = useSelector((state: AppRootState) => state.LoginPage.user)
  const { userId: receiverId } = useParams()

  useEffect(() => {
    if (me && receiverId) {
      dispatch(getConversationRequest(me.userId, Number(receiverId)))
    }
  }, [receiverId])

  return (
    <div className={conversation}>
      <div className={conversationHistory}>
        {messages.map((message) => (
          <ChatMessageItem key={message.id} sender={message.senderId} text={message.text} />
        ))}
      </div>
      <div className={newMessageForm}>
        <textarea className={newMessageField} placeholder="Write your message..." />
        <button className={sendMessageButton}>Send message</button>
      </div>
    </div>
  )
}
