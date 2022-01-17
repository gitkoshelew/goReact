import s from './Login.module.scss'
import { Form, Formik } from 'formik'
import { authenticationSchema } from './validations/LoginValidation'
import { TextField } from '../../components/TextField/TextField'
import { AppRootStateType, useAppDispatch } from '../../../bll/store/store'
import { LoginRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import Preloader from '../../components/preloader/preloader'
import { Home } from '../Home/Home'
import { LoginErrorMsg } from '../../components/ErrorMsgLogin/LoginErrorMsg'
import { UserRequestData } from '../../../dal/api_client/AuthService'
import { closedEye, openedEye } from '../../svgWrapper/LoginSvgWrapper'
import { useState } from 'react'
import { NavLink } from 'react-router-dom'
import { PATH } from '../../Routes/RoutesInfo'

const {
  authenticationForm,
  authenticationTitle,
  sendReqBtn,
  sendReqErrorBtn,
  passwordField,
  passwordRenderEye,
  signUpLink,
  loginBtnLinkBlock,
} = s

export const LoginPage = () => {
  const dispatch = useAppDispatch()
  const LoginPageLoadingStatus = useSelector((state: AppRootStateType) => state.LoginPage.loadingStatus)
  const userProfile = useSelector((state: AppRootStateType) => state.LoginPage.user)
  const ErrorMsg = useSelector((state: AppRootStateType) => state.LoginPage.errorMsg)

  const [isPasswordOpen, setIsPasswordOpen] = useState(false)

  const showPasswordHandler = () => {
    setIsPasswordOpen((currentValue) => !currentValue)
  }

  const errMsg = ErrorMsg && <LoginErrorMsg ErrorMsg={ErrorMsg} />
  const correctEyeRender = isPasswordOpen ? closedEye : openedEye
  const correctPasswordInputType = isPasswordOpen ? 'text' : 'password'

  if (LoginPageLoadingStatus === 'loading') {
    return <Preloader />
  }
  if (userProfile) {
    return <Home />
  }

  return (
    <Formik
      initialValues={{
        email: '',
        password: '',
      }}
      validationSchema={authenticationSchema}
      onSubmit={(user: UserRequestData) => {
        dispatch(LoginRequest(user))
      }}
    >
      {(formik) => (
        <Form>
          <div className={authenticationForm}>
            <div className={authenticationTitle}>SING IN</div>
            <TextField inputType={'login'} label="Email" name="email" type="text" />
            <div className={passwordField}>
              <TextField inputType={'register'} label="Password" name="password" type={correctPasswordInputType} />
              <img onClick={showPasswordHandler} className={passwordRenderEye} src={correctEyeRender} alt="eye" />
            </div>
            <div className={loginBtnLinkBlock}>
              <button
                className={formik.errors.email || formik.errors.password ? sendReqErrorBtn : sendReqBtn}
                type="submit"
              >
                Sign In
              </button>
              <NavLink to={PATH.REGISTRATION} className={signUpLink}>
                Sign Up
              </NavLink>
            </div>
          </div>
          {errMsg}
        </Form>
      )}
    </Formik>
  )
}
