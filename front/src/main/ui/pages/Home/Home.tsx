import React, { useEffect } from 'react'
import { FirstPage } from './firstPage/FirstPage'
import { SecondPage } from './SecondPage/SecondPage'
import { ThirdPage } from './ThirdPage/ThirdPage'
import { FourthPage } from './FourthPage/FourthPage'
import { FifthPage } from './FifthPage/FifthPage'
import { SixthPage } from './SixthPage/SixthPage'
import { AppRootStateType, useAppDispatch } from '../../../bll/store/store'
import { MeRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import { LogInResponseType } from '../../../dal/api_client/AuthService'

export const Home = () => {
  const dispatch = useAppDispatch()
  const userProfile = useSelector<AppRootStateType, LogInResponseType | null>((state) => state.LoginPage.user)
  useEffect(() => {
    debugger
    if (localStorage.getItem('token') && userProfile === null) {
      dispatch(MeRequest())
    }
  }, [])

  return (
    <div>
      <FirstPage />
      <SecondPage />
      <ThirdPage />
      <FourthPage />
      <FifthPage />
      <SixthPage />
    </div>
  )
}
