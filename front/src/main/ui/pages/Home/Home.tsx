import React, { useEffect } from 'react'
import { FirstPage } from './WelcomePage/FirstPage'
import { SecondPage } from './CeoMsgPage/SecondPage'
import { ThirdPage } from './FavoriteRoomsPage/ThirdPage'
import { FourthPage } from './ServicesPage/FourthPage'
import { FifthPage } from './PlacesNearbyPage/FifthPage'
import { SixthPage } from './UsersFeedbackPage/SixthPage'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { MeRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'

export const Home = () => {
  const dispatch = useAppDispatch()
  const userProfile = useSelector((state: AppRootState) => state.LoginPage.user)

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
