import { Home } from '../pages/Home/Home'
import { Navigate, Route, Routes } from 'react-router-dom'
import { Hotels } from '../pages/Hotels/Hotels'
import { AboutUs } from '../pages/AboutUs/AboutUs'
import { Error404 } from '../pages/error404/error404'
import { LoginPage } from '../pages/Login/LoginPage'
import { Gallery } from '../pages/Gallery/Gallery'
import { Room } from '../pages/Room/Room'
import { Booking } from '../pages/Booking/Booking'
import { Service } from '../pages/Service/Service'
import { Basket } from '../pages/Basket/Basket'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../bll/store/store'
import { LogInResponse } from '../../dal/api_client/AuthService'
import { ReactJSXElement } from '@emotion/react/types/jsx-namespace'
import { RegistrationPage } from '../pages/Registration/RegistrationPage'
import { RegisterPageLoadingStatus } from '../../bll/reducers/RegistrationPageReducer/registrationPage-reducer'
import { Chat } from '../pages/Chat/Chat'
import { Conversation } from '../components/Conversation/Conversation'
import { LoadingStatuses } from '../../bll/reducers/types/enum'

type LoginWrapperType = {
  children: ReactJSXElement
  user: LogInResponse | null
}
type RegisterWrapper = {
  children: ReactJSXElement
  loadingStatus: RegisterPageLoadingStatus
}

export const PATH = {
  HOME: '/home',
  LOGIN: '/login',
  HOTELS: '/hotels',
  ABOUT_US: '/aboutus',
  ROOM: '/room',
  SERVICE: '/service',
  BOOKING: '/booking',
  GALLERY: '/gallery',
  BASKET: '/basket',
  REGISTRATION: '/registration',
  CHAT: '/chat/*',
  CHAT_CONVERSATION: '/:userId',
}

export const RoutesInfo = () => {
  const userProfile = useSelector((state: AppRootState) => state.LoginPage.user)

  const registerPageLoadingStatus = useSelector((state: AppRootState) => state.RegisterPage.loadingStatus)

  const LoginWrapper = ({ children, user }: LoginWrapperType) => {
    return user ? <Navigate to={PATH.HOME} replace /> : children
  }
  const RegisterWrapper = ({ children, loadingStatus }: RegisterWrapper) => {
    return loadingStatus === LoadingStatuses.SUCCESS ? <Navigate to={PATH.LOGIN} replace /> : children
  }
  return (
    <div>
      <Routes>
        <Route path={'/'} element={<Navigate replace to={PATH.HOME} />} />
        <Route path={PATH.HOME} element={<Home />} />
        <Route path={PATH.HOTELS} element={<Hotels />} />
        <Route path={PATH.ABOUT_US} element={<AboutUs />} />
        <Route
          path={PATH.LOGIN}
          element={
            <LoginWrapper user={userProfile}>
              <LoginPage />
            </LoginWrapper>
          }
        />
        <Route
          path={PATH.REGISTRATION}
          element={
            <RegisterWrapper loadingStatus={registerPageLoadingStatus}>
              <RegistrationPage />
            </RegisterWrapper>
          }
        />
        <Route path={PATH.GALLERY} element={<Gallery />} />
        <Route path={PATH.ROOM} element={<Room />} />
        <Route path={PATH.BOOKING} element={<Booking />} />
        <Route path={PATH.SERVICE} element={<Service />} />
        <Route path={PATH.BASKET} element={<Basket />} />
        <Route
          path={PATH.CHAT}
          element={
            <Chat>
              <Routes>
                <Route path={PATH.CHAT_CONVERSATION} element={<Conversation />} />
              </Routes>
            </Chat>
          }
        />

        <Route path={'*'} element={<Error404 />} />
      </Routes>
    </div>
  )
}
