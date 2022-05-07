import React, { useState } from 'react'
import s from './userNavBarView.module.scss'
import { AppRootState, useAppDispatch } from '../../../../bll/store/store'
import { LogOutRequest } from '../../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import Preloader from '../../../components/preloader/preloader'
import { LogInResponse } from '../../../../dal/api_client/AuthService'
import { LoadingStatuses } from '../../../../bll/reducers/types/enum'

const { userNav, userPhoto, userName, logoutField, userNavView, preloaderNav } = s

export type UserNavBarViewPropsType = {
  user: LogInResponse
}

export const UserNavBarView = ({ user }: UserNavBarViewPropsType) => {
  const dispatch = useAppDispatch()
  const loadingStatus = useSelector((state: AppRootState) => state.LoginPage.loadingStatus)

  const [isLogoutVisible, setIsLogoutVisible] = useState(false)

  const onLogoutFieldView = () => {
    setIsLogoutVisible((currentValue) => !currentValue)
  }

  const handleLogout = () => {
    dispatch(LogOutRequest())
    setIsLogoutVisible(false)
  }

  if (loadingStatus === LoadingStatuses.LOADING) {
    return (
      <div className={preloaderNav}>
        <Preloader />
      </div>
    )
  }

  return (
    <div className={userNavView}>
      <div onClick={onLogoutFieldView} className={userNav}>
        <div className={userPhoto}>
          <img src={user.photo} alt="userPhoto" />
        </div>
        <div className={userName}>{`${user.name} ${user.sName}`}</div>
      </div>
      {isLogoutVisible && (
        <div onClick={handleLogout} className={logoutField}>
          Logout
        </div>
      )}
    </div>
  )
}
