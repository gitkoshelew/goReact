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
import { ChatPageReducer } from '../reducers/ChatPageReducer/chatPage-reducer'
import {
  closeChannelSagaWorker,
  fetchInitMessagesSagaWorker,
  fetchUsersSagaWorker,
  getConversationSagaWorker,
  sendMessageSagaWorker,
  setConversationOpenedSagaWorker,
} from '../reducers/ChatPageReducer/chatPage-saga'
import { openChannelSagaWorker } from '../reducers/ChatPageReducer/socketChannel'

const sagaMiddleware = createSagaMiddleware()

const rootReducer = combineReducers({
  BookingRegForm: BookingRegFormReducer,
  BookingRoomPick: BookingRoomPickReducer,
  LoginPage: LoginPageReducer,
  RegisterPage: RegisterPageReducer,
  ChatPage: ChatPageReducer,
})

export type RootReducer = typeof rootReducer

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // to fix error-warning in log at upload pet img action
    }).prepend(sagaMiddleware),
})

export type AppRootState = ReturnType<typeof rootReducer>

export type AppDispatch = typeof store.dispatch

export const useAppDispatch = () => useDispatch<AppDispatch>()

//sagaWatcher
sagaMiddleware.run(rootWatcher)

function* rootWatcher() {
  yield takeEvery('BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD', BookingUploadPetImgSagaWorker)
  yield takeEvery('BOOKING_ROOM_PICK/NEW_IS_RENT_ROOMS_FOR_CALENDAR', BookingRoomPickSagaWorker)
  yield takeEvery('LOGIN_PAGE/LOGIN_SAGA', LoginPageLoginSagaWorker)
  yield takeEvery('LOGIN_PAGE/LOGOUT_SAGA', LoginPageLogoutSagaWorker)
  yield takeEvery('LOGIN_PAGE/ME_SAGA', LoginPageMeRequestSagaWorker)
  yield takeEvery('REGISTER_PAGE/REGISTER_SAGA', RegisterPageSagaWorker)
  yield takeEvery('CHAT_PAGE/FETCH_USERS_SAGA', fetchUsersSagaWorker)
  yield takeEvery('CHAT_PAGE/FETCH_INIT_MESSAGES_SAGA', fetchInitMessagesSagaWorker)
  yield takeEvery('CHAT_PAGE/GET_CONVERSATION_SAGA', getConversationSagaWorker)
  yield takeEvery('CHAT_PAGE/OPEN_CHANNEL', openChannelSagaWorker)
  yield takeEvery('CHAT_PAGE/CLOSE_CHANNEL', closeChannelSagaWorker)
  yield takeEvery('CHAT_PAGE/USER_SEND_MESSAGE', sendMessageSagaWorker)
  yield takeEvery('CHAT_PAGE/SET_CONVERSATION_OPENED', setConversationOpenedSagaWorker)
}
