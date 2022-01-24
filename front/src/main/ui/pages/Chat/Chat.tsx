import React, { FC, useEffect, useMemo } from 'react'
import s from './Chat.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { fetchUsersRequest } from '../../../bll/reducers/ChatPageReducer/chatPage-saga'
import { UserChatItem } from '../../components/UserChatItem/UserChatItem'
import { MeRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'

const { usersList, chat } = s

export const Chat: FC = ({ children }) => {
  const dispatch = useAppDispatch()
  const me = useSelector((state: AppRootState) => state.LoginPage.user)
  const allUsers = useSelector((state: AppRootState) => state.ChatPage.users)

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
      <div className={usersList}>{chatPartners}</div>
      {children}
    </div>
  )
}
