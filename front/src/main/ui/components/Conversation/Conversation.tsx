import React, { ChangeEvent, useEffect, useState } from 'react'
import s from './Conversation.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { ChatMessageItem } from '../ChatMessageItem/ChatMessageItem'
import { useParams } from 'react-router-dom'
import {
  addMessageRequest,
  closeChannelRequest,
  getConversationRequest,
} from '../../../bll/reducers/ChatPageReducer/chatPage-saga'
import { openChannelRequest } from '../../../bll/reducers/ChatPageReducer/socketChannel'
import { MeRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'

const { conversation, conversationHistory, newMessageForm, newMessageField, sendMessageButton } = s

export const Conversation = () => {
  const [messageText, setMessageText] = useState<string>('')
  const dispatch = useAppDispatch()
  const messages = useSelector((state: AppRootState) => state.ChatPage.messages)
  const me = useSelector((state: AppRootState) => state.LoginPage.user)
  const currentConversation = useSelector((state: AppRootState) => state.ChatPage.conversationId)
  const socketChannel = useSelector((state: AppRootState) => state.ChatPage.socketChannel)
  const { userId: consumerId } = useParams()

  useEffect(() => {
    if ((me === null && localStorage.getItem('token')) || localStorage.getItem('MockToken')) {
      dispatch(MeRequest())
    }
  }, [])

  useEffect(() => {
    if (me && consumerId) {
      dispatch(getConversationRequest(me.userId, Number(consumerId)))
      dispatch(openChannelRequest(me.userId, Number(consumerId)))
    }

    return () => {
      if (socketChannel) {
        dispatch(closeChannelRequest())
      }
    }
  }, [consumerId, me])

  const handleChangeMessage = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setMessageText(e.currentTarget.value)
  }
  const handleSendMessage = () => {
    if (me && consumerId && currentConversation) {
      const newMessage = {
        producerId: me.userId,
        consumerId: Number(consumerId),
        conversationId: currentConversation,
        text: messageText,
      }
      if (socketChannel) {
        dispatch(addMessageRequest(socketChannel, newMessage))
      }
    }
    setMessageText('')
  }

  return (
    <div className={conversation}>
      <div className={conversationHistory}>
        {messages.map((message) => (
          <ChatMessageItem key={message.id} sender={message.producerId} text={message.text} />
        ))}
      </div>
      <div className={newMessageForm}>
        <textarea
          className={newMessageField}
          placeholder="Write your message..."
          value={messageText}
          onChange={handleChangeMessage}
        />
        <button className={sendMessageButton} onClick={handleSendMessage}>
          Send message
        </button>
      </div>
    </div>
  )
}
