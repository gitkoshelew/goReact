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

export const Home = () => {
  const dispatch = useAppDispatch()
  const userProfile = useSelector((state: AppRootStateType) => state.LoginPage.user)

  useEffect(() => {
    if ((userProfile === null && localStorage.getItem('token')) || localStorage.getItem('MockToken')) {
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
