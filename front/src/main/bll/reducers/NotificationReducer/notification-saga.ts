import { apply, call, put, select } from 'redux-saga/effects'
import { addNotification, removeNotification, setNotificationSocketChannel } from './notification-reducer'
import { Socket } from 'socket.io-client'
import { AppRootState } from '../../store/store'
import { NotificationRaw } from './socketChannel'

export function* addNotificationSagaWorker(action: ReturnType<typeof addNotificationRequest>) {
  yield put(addNotification(action.notification))
}

export const addNotificationRequest = (notification: NotificationRaw) => ({
  type: 'NOTIFICATION/ADD_NOTIFICATION',
  notification,
})

export function* removeNotificationSagaWorker(action: ReturnType<typeof removeNotificationRequest>) {
  yield put(removeNotification(action.notificationId))
}

export const removeNotificationRequest = (notificationId: string) => ({
  type: 'NOTIFICATION/REMOVE_NOTIFICATION',
  notificationId,
})

export function* confirmNotificationSagaWorker(action: ReturnType<typeof confirmNotificationRequest>) {
  const socket: Socket = yield select((state: AppRootState) => state.Notification.socketChannel)
  const notificationToConfirm: NotificationRaw = yield select((state: AppRootState) =>
    state.Notification.notifications.find((notification) => notification.content.id === action.notificationId)
  )
  yield apply(socket, socket.emit, ['CLIENT_RECEIVED_NOTIFICATION', notificationToConfirm])
}

export const confirmNotificationRequest = (notificationId: string) => ({
  type: 'NOTIFICATION/CONFIRM_NOTIFICATION',
  notificationId,
})

export function* closeNotificationChannelSagaWorker() {
  const socket: Socket = yield select((state: AppRootState) => state.Notification.socketChannel)
  yield call([socket, socket.disconnect])
  yield put(setNotificationSocketChannel({ socketChannel: null }))
}

export const closeNotificationChannelRequest = () => ({
  type: 'NOTIFICATION/CLOSE_CHANNEL',
})
