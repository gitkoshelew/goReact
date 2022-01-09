import { userPhotoBoilerPlate } from '../../../svgWrapper/navBarSvgWrapper'
import React, { useState } from 'react'
import { LogInResponseType } from '../../../../dal/api_client/API'
import s from './userNavBarView.module.scss'
import { AppRootStateType, useAppDispatch } from '../../../../bll/store/store'
import { LogOutRequest } from '../../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import { LoginPageLoadingStatusType } from '../../../../bll/reducers/LoginPageReduser/loginPage-reducer'
import Preloader from '../../../components/preloader/preloader'

const { userNav, userPhoto, userName, logoutField, userNavView, preloaderNav } = s

export type UserNavBarViewPropsType = {
  user: LogInResponseType
}

export const UserNavBarView = ({ user }: UserNavBarViewPropsType) => {
  const dispatch = useAppDispatch()
  const loadingStatus = useSelector<AppRootStateType, LoginPageLoadingStatusType>(
    (state) => state.LoginPage.loadingStatus
  )

  const [isLogoutVisible, setIsLogoutVisible] = useState(false)

  const onLogoutFieldView = () => {
    setIsLogoutVisible(!isLogoutVisible)
  }

  const onLogoutHandler = () => {
    dispatch(LogOutRequest())
    setIsLogoutVisible(false)
  }

  if (loadingStatus === 'loading') {
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
          {user.photo === 'PhotoURL...' ? (
            <img src={userPhotoBoilerPlate} alt="userBoilerPlate" />
          ) : (
            <img src={user.photo} alt="userPhoto" />
          )}
        </div>
        <div className={userName}>{`${user.name} ${user.sName}`}</div>
      </div>
      {isLogoutVisible && (
        <div onClick={onLogoutHandler} className={logoutField}>
          Logout
        </div>
      )}
    </div>
  )
}
