import s from './RegistrationPage.module.scss'
import { Form, Formik } from 'formik'
import { TextField } from '../../components/TextField/TextField'
import { AppRootStateType } from '../../../bll/store/store'
import { useSelector } from 'react-redux'
import { LoginPageLoadingStatusType } from '../../../bll/reducers/LoginPageReduser/loginPage-reducer'
import Preloader from '../../components/preloader/preloader'
import { Home } from '../Home/Home'
import { LoginErrorMsg } from '../../components/ErrorMsgLogin/LoginErrorMsg'
import { LogInResponse, RegisterResponse } from '../../../dal/api_client/AuthService'
import { closedEye, openedEye } from '../../svgWrapper/LoginSwgWrapper'
import { useState } from 'react'
import { PATH } from '../../Routes/RoutesInfo'
import { NavLink } from 'react-router-dom'
import { RegistrationSchema } from './validations/RegisterValidation'

type OnSubmitValues = {
  email: string
  password: string
  name: string
  sName: string
  mName: string
  sex: string
  birthDate: string
  address: string
  phone: string
}

const {
  RegistrationForm,
  RegistrationTitle,
  sendReqBtn,
  sendReqErrorBtn,
  passwordField,
  passwordRenderEye,
  backLink,
  loginBtnLinkBlock,
  sexRadioGroup,
  registerRadioGroup,
} = s

export const RegistrationPage = () => {
  const LoginPageLoadingStatus = useSelector<AppRootStateType, LoginPageLoadingStatusType>(
    (state) => state.LoginPage.loadingStatus
  )
  const userProfile = useSelector<AppRootStateType, LogInResponse | null>((state) => state.LoginPage.user)
  const ErrorMsg = useSelector<AppRootStateType, string>((state) => state.LoginPage.errorMsg)

  const [isPasswordOpen, setIsPasswordOpen] = useState<boolean>(false)

  const showPasswordHandler = () => {
    setIsPasswordOpen(!isPasswordOpen)
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
        name: '',
        sName: '',
        mName: '',
        sex: '',
        birthDate: '',
        address: '',
        phone: '',
      }}
      validationSchema={RegistrationSchema}
      onSubmit={(newUser: OnSubmitValues) => {
        const dataForRequest: RegisterResponse = {
          ...newUser,
          photo: 'PhotoURL...',
          verified: true,
          role: 'client',
        }
        console.log(dataForRequest)
      }}
    >
      {(formik) => (
        <Form>
          <div className={RegistrationForm}>
            <div className={RegistrationTitle}>SIGN UP</div>
            <TextField inputMsgLabel={'Email'} inputType={'register'} name="email" type="text" />
            <div className={passwordField}>
              <TextField
                inputMsgLabel={'Password'}
                inputType={'register'}
                name="password"
                type={correctPasswordInputType}
              />
              <img onClick={showPasswordHandler} className={passwordRenderEye} src={correctEyeRender} alt="eye" />
            </div>
            <TextField inputMsgLabel={'Name'} inputType={'register'} name="name" type="text" />
            <TextField inputMsgLabel={'Second name'} inputType={'register'} name="sName" type="text" />
            <TextField inputMsgLabel={'Middle name'} inputType={'register'} name="mName" type="text" />
            <TextField inputMsgLabel={'Address'} inputType={'register'} name="address" type="text" />
            <TextField inputMsgLabel={'Phone'} inputType={'register'} name="phone" type="text" />
            <div className={registerRadioGroup}>
              <div>Sex</div>
              <div className={sexRadioGroup}>
                <TextField inputMsgLabel={'Male'} inputType={'register'} name="sex" type="radio" value={'male'} />
                <TextField inputMsgLabel={'Female'} inputType={'register'} name="sex" type="radio" value={'female'} />
              </div>
            </div>
            <TextField inputMsgLabel={'Birthday date'} inputType={'register'} name="birthDate" type="date" />

            <div className={loginBtnLinkBlock}>
              <button
                className={formik.errors.email || formik.errors.password ? sendReqErrorBtn : sendReqBtn}
                type="submit"
              >
                Sign UP
              </button>
              <NavLink to={PATH.LOGIN} className={backLink}>
                Sign IN
              </NavLink>
            </div>
          </div>
          {errMsg}
        </Form>
      )}
    </Formik>
  )
}
