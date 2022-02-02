import { delay, put } from 'redux-saga/effects'
import { NotificationData, setNotificationData, toggleNotification } from './notification-reducer'

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
