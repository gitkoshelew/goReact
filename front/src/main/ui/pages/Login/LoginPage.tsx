import s from './Login.module.scss'
import { Form, Formik } from 'formik'
import { authenticationSchema } from './validations/authValidation'
import { TextField } from '../../components/TextField/TextField'
import { LogInResponseType, UserRequestDataType } from '../../../dal/api_client/API'
import { AppRootStateType, useAppDispatch } from '../../../bll/store/store'
import { LoginRequest } from '../../../bll/reducers/LoginPageReduser/loginPage-saga'
import { useSelector } from 'react-redux'
import { LoginPageLoadingStatusType } from '../../../bll/reducers/LoginPageReduser/loginPage-reducer'
import Preloader from '../../components/preloader/preloader'
import { Home } from '../Home/Home'
import { LoginErrorMsg } from '../../components/ErrorMsgLogin/LoginErrorMsg'

const { authenticationForm, authenticationTitle, sendReqBtn, sendReqErrorBtn } = s

export const LoginPage = () => {
  const dispatch = useAppDispatch()
  const LoginPageLoadingStatus = useSelector<AppRootStateType, LoginPageLoadingStatusType>(
    (state) => state.LoginPage.loadingStatus
  )
  const userProfile = useSelector<AppRootStateType, LogInResponseType | null>((state) => state.LoginPage.user)
  const ErrorMsg = useSelector<AppRootStateType, string>((state) => state.LoginPage.errorMsg)

  const errMsg = ErrorMsg && <LoginErrorMsg ErrorMsg={ErrorMsg} />

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
      onSubmit={(user: UserRequestDataType) => {
        dispatch(LoginRequest(user))
      }}
    >
      {(formik) => (
        <Form>
          <div className={authenticationForm}>
            <div className={authenticationTitle}>Log in</div>
            <TextField label="Email" name="email" type="text" />
            <TextField label="Password" name="password" type="text" />
            <button
              className={formik.errors.email || formik.errors.password ? sendReqErrorBtn : sendReqBtn}
              type="submit"
            >
              Login
            </button>
          </div>
          {errMsg}
        </Form>
      )}
    </Formik>
  )
}
