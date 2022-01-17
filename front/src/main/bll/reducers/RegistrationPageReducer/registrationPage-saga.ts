import { call, put } from 'redux-saga/effects'
import { AuthAPI, RegisterRequestUser } from '../../../dal/api_client/AuthService'
import { reqRegisterError, reqRegisterStart, reqRegisterSuccess } from './registrationPage-reducer'

export function* RegisterPageSagaWorker(action: RegistrationRequest) {
  try {
    yield put(reqRegisterStart())
    yield call(AuthAPI.RegisterAPI, action.user)
    yield put(reqRegisterSuccess())
  } catch (err) {
    if (err instanceof Error) {
      yield put(reqRegisterError({ errorMsg: err.message }))
    }
  }
}

export const RegisterRequest = (user: RegisterRequestUser) => ({
  type: 'REGISTER_PAGE/REGISTER_SAGA',
  user,
})

type RegistrationRequest = ReturnType<typeof RegisterRequest>
