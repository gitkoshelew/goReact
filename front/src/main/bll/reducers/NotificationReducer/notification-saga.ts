import { call, delay, put, select } from 'redux-saga/effects'
import {
  NotificationData,
  setNotificationData,
  setNotificationSocketChannel,
  toggleNotification,
} from './notification-reducer'
import { Socket } from 'socket.io-client'
import { AppRootState } from '../../store/store'

export function* showNotificationSagaWorker(action: ReturnType<typeof showNotificationRequest>) {
  yield put(setNotificationData(action.data))
  yield put(toggleNotification({ isOpen: true }))
  yield delay(6000)
  yield put(toggleNotification({ isOpen: false }))
}

export const showNotificationRequest = (notificationData: NotificationData) => ({
  type: 'NOTIFICATION/SHOW_NOTIFICATION',
  data: notificationData,
})

export function* closeNotificationChannelSagaWorker() {
  const socket: Socket = yield select((state: AppRootState) => state.Notification.socketChannel)
  yield call([socket, socket.disconnect])
  yield put(setNotificationSocketChannel({ socketChannel: null }))
}

export const closeNotificationChannelRequest = () => ({
  type: 'NOTIFICATION/CLOSE_CHANNEL',
})
