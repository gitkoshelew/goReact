import s from './Login.module.scss'
import { Form, Formik } from 'formik'
import { authenticationSchema } from './validations/LoginValidation'
import { TextField } from '../../components/TextField/TextField'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { LoginRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import Preloader from '../../components/preloader/preloader'
import { Home } from '../Home/Home'
import { LoginErrorMsg } from '../../components/ErrorMsgLogin/LoginErrorMsg'
import { UserRequestData } from '../../../dal/api_client/AuthService'
import { closedEye, Github, Linkedin, openedEye } from '../../svgWrapper/LoginSvgWrapper'
import { NavLink } from 'react-router-dom'
import { PATH } from '../../Routes/RoutesInfo'
import { useState } from 'react'
import { LoadingStatuses } from '../../../bll/reducers/types/enum'

const {
  authenticationForm,
  authenticationTitle,
  sendReqBtn,
  sendReqErrorBtn,
  passwordField,
  passwordRenderEye,
  signUpLink,
  loginBtnLinkBlock,
  socialBtnWrapper,
  loginBtn,
  loginBtnGithub,
  loginBtnLinkedin,
  icon,
  linkSocialMedia,
} = s

export const LoginPage = () => {
  const dispatch = useAppDispatch()
  const LoginPageLoadingStatus = useSelector((state: AppRootState) => state.LoginPage.loadingStatus)
  const userProfile = useSelector((state: AppRootState) => state.LoginPage.user)
  const ErrorMsg = useSelector((state: AppRootState) => state.LoginPage.errorMsg)

  const [isPasswordOpen, setIsPasswordOpen] = useState(false)

  const showPasswordHandler = () => {
    setIsPasswordOpen((currentValue) => !currentValue)
  }

  const errMsg = ErrorMsg && <LoginErrorMsg ErrorMsg={ErrorMsg} />
  const correctEyeRender = isPasswordOpen ? closedEye : openedEye
  const correctPasswordInputType = isPasswordOpen ? 'text' : 'password'

  const githubClassName = ` ${loginBtn} ${loginBtnGithub}`
  const linkedinClassName = ` ${loginBtn} ${loginBtnLinkedin}`

  const githubUrl = process.env.REACT_APP_API_SOCIAL_AUTH_LINK + 'api/gitlogin'

  const linkedinUrl = process.env.REACT_APP_API_SOCIAL_AUTH_LINK + 'api/linkedinlogin'

  if (LoginPageLoadingStatus === LoadingStatuses.LOADING) {
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
            <div className={socialBtnWrapper}>
              <div className={githubClassName}>
                <img src={Github} alt="" className={icon} />
                <a className={linkSocialMedia} href={githubUrl}>
                  Github
                </a>
              </div>
              <div className={linkedinClassName}>
                <img src={Linkedin} alt="" className={icon} />
                <a className={linkSocialMedia} href={linkedinUrl}>
                  Linkedin
                </a>
              </div>
            </div>
          </div>
          {errMsg}
        </Form>
      )}
    </Formik>
  )
}
