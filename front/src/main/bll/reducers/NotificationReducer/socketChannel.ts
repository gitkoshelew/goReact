import { io, Socket } from 'socket.io-client'
import { eventChannel } from 'redux-saga'
import { call, put, take } from 'redux-saga/effects'
import { NotificationData, setNotificationSocketChannel } from './notification-reducer'
import { addNotificationRequest } from './notification-saga'

export type NotificationRaw = {
  content: NotificationData
  fields: any
  properties: any
}

function createSocketConnections(clientId: number) {
  return io('ws://localhost:5000', { query: { clientId } })
}

const createEventProviderChannel = (eventProvider: Socket) => {
  return eventChannel((emit) => {
    const receivedNotificationHandler = (notification: string) => {
      emit(JSON.parse(notification))
    }
    eventProvider.on('BROKER_RECEIVED_NOTIFICATION', receivedNotificationHandler)

    return () => {
      eventProvider.off('BROKER_RECEIVED_NOTIFICATION', receivedNotificationHandler)
    }
  })
}

export function* openNotificationChannelSagaWorker(action: ReturnType<typeof openNotificationChannelRequest>) {
  const { clientId } = action.payload
  const eventProvider: Socket = yield call(createSocketConnections, clientId)
  yield put(setNotificationSocketChannel({ socketChannel: eventProvider }))

  // @ts-ignore
  const eventProviderChannel = yield call(createEventProviderChannel, eventProvider)

  try {
    while (true) {
      const payload: NotificationRaw = yield take(eventProviderChannel)
      yield put(addNotificationRequest(payload))
    }
  } catch (err) {
    console.log(err)
  }
}

export const openNotificationChannelRequest = (clientId: number) => ({
  type: 'NOTIFICATION/OPEN_CHANNEL',
  payload: {
    clientId,
  },
})
