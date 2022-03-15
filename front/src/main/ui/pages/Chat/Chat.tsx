import React, { FC, useEffect, useMemo } from 'react'
import s from './Chat.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { fetchUsersRequest } from '../../../bll/reducers/ChatPageReducer/chatPage-saga'
import { UserChatItem } from '../../components/UserChatItem/UserChatItem'
import { MeRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import goBackIcon from '../../../../assets/img/chat/arrowBack.png'
import { useNavigate } from 'react-router-dom'

const { usersList, chat, showConversationsButton, hiddenUserList, hiddenConversationButton } = s

export const Chat: FC = ({ children }) => {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()
  const me = useSelector((state: AppRootState) => state.LoginPage.user)
  const allUsers = useSelector((state: AppRootState) => state.ChatPage.users)
  const isConversationOpened = useSelector((state: AppRootState) => state.ChatPage.isConversationOpened)

  useEffect(() => {
    if ((me === null && localStorage.getItem('token')) || localStorage.getItem('MockToken')) {
      dispatch(MeRequest())
    }
  }, [])

  useEffect(() => {
    dispatch(fetchUsersRequest())
  }, [])

  const chatPartners = useMemo(() => {
    return allUsers
      .filter((user) => {
        return me ? user.userId !== me.userId : user
      })
      .map((user) => (
        <UserChatItem
          key={user.userId}
          userId={user.userId}
          name={user.name}
          sName={user.sName}
          role={user.role}
          photo={user.photo}
        />
      ))
  }, [allUsers])

  return (
    <div className={chat}>
      <div
        className={`${showConversationsButton} ${!isConversationOpened ? hiddenConversationButton : ''}`}
        onClick={() => navigate('/chat')}
      >
        <img src={goBackIcon} alt="Show conversations" />
      </div>
      <div className={`${usersList} ${isConversationOpened ? hiddenUserList : ''}`}>{chatPartners}</div>
      {children}
    </div>
  )
}
