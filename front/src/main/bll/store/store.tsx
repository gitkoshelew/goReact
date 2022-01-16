import { combineReducers } from 'redux'
import { configureStore } from '@reduxjs/toolkit'
import createSagaMiddleware from 'redux-saga'
import { BookingRegFormReducer } from '../reducers/BookingRegFormReducer/BookingRegForm-reducer'
import { useDispatch } from 'react-redux'
import { takeEvery } from 'redux-saga/effects'
import { BookingUploadPetImgSagaWorker } from '../reducers/BookingRegFormReducer/BookindRegForm-saga'
import { BookingRoomPickReducer } from '../reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'
import { BookingRoomPickSagaWorker } from '../reducers/BookingRoomsPickReducer/BookingRoomPick-saga'
import { LoginPageReducer } from '../reducers/LoginPageReduser/loginPage-reducer'
import {
  LoginPageLoginSagaWorker,
  LoginPageLogoutSagaWorker,
  LoginPageMeRequestSagaWorker,
} from '../reducers/LoginPageReduser/loginPage-saga'
import { RegisterPageReducer } from '../reducers/RegistrationPageReducer/registrationPage-reducer'
import { RegisterPageSagaWorker } from '../reducers/RegistrationPageReducer/registrationPage-saga'

const sagaMiddleware = createSagaMiddleware()

const rootReducer = combineReducers({
  BookingRegForm: BookingRegFormReducer,
  BookingRoomPick: BookingRoomPickReducer,
  LoginPage: LoginPageReducer,
  RegisterPage: RegisterPageReducer,
})

export type RootReducerType = typeof rootReducer

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // to fix error-warning in log at upload pet img action
    }).prepend(sagaMiddleware),
})

export type AppRootStateType = ReturnType<typeof rootReducer>

export type AppDispatchType = typeof store.dispatch

export const useAppDispatch = () => useDispatch<AppDispatchType>()

//sagaWatcher
sagaMiddleware.run(rootWatcher)

function* rootWatcher() {
  yield takeEvery('BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD', BookingUploadPetImgSagaWorker)
  yield takeEvery('BOOKING_ROOM_PICK/NEW_IS_RENT_ROOMS_FOR_CALENDAR', BookingRoomPickSagaWorker)
  yield takeEvery('LOGIN_PAGE/LOGIN_SAGA', LoginPageLoginSagaWorker)
  yield takeEvery('LOGIN_PAGE/LOGOUT_SAGA', LoginPageLogoutSagaWorker)
  yield takeEvery('LOGIN_PAGE/ME_SAGA', LoginPageMeRequestSagaWorker)
  yield takeEvery('REGISTER_PAGE/REGISTER_SAGA', RegisterPageSagaWorker)
}
