import { call, delay, put, select } from 'redux-saga/effects'
import {
  addNotificationToQueue,
  NotificationData,
  removeNotificationFromQueue,
  setCurrentNotification,
  setNotificationSocketChannel,
  toggleNotification,
} from './notification-reducer'
import { Socket } from 'socket.io-client'
import { AppRootState } from '../../store/store'

export function* addNotificationToQueueSagaWorker(action: ReturnType<typeof addNotificationToQueueRequest>) {
  yield put(addNotificationToQueue(action.data))
  yield put(showNotificationRequest())
}

export const addNotificationToQueueRequest = (notificationData: NotificationData) => ({
  type: 'NOTIFICATION/ADD_NOTIFICATION',
  data: notificationData,
})

export function* removeNotificationFromQueueSagaWorker() {
  yield put(removeNotificationFromQueue())
}

export const removeNotificationFromQueueRequest = () => ({
  type: 'NOTIFICATION/REMOVE_NOTIFICATION',
})

export function* showNotificationSagaWorker() {
  const nextNotificationToShow: NotificationData = yield select(
    (state: AppRootState) => state.Notification.allNotifications[0]
  )
  yield put(setCurrentNotification(nextNotificationToShow))
  yield put(toggleNotification({ isOpen: true }))
  yield delay(6000)
  yield put(toggleNotification({ isOpen: false }))
  yield put(removeNotificationFromQueueRequest())
}

export const showNotificationRequest = () => ({
  type: 'NOTIFICATION/SHOW_NOTIFICATION',
})

export function* closeNotificationChannelSagaWorker() {
  const socket: Socket = yield select((state: AppRootState) => state.Notification.socketChannel)
  yield call([socket, socket.disconnect])
  yield put(setNotificationSocketChannel({ socketChannel: null }))
}

export const closeNotificationChannelRequest = () => ({
  type: 'NOTIFICATION/CLOSE_CHANNEL',
})
