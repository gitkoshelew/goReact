import { io, Socket } from 'socket.io-client'
import { call, put, take } from 'redux-saga/effects'
import { eventChannel } from 'redux-saga'
import { addMessage, fetchError, setSocketChannel } from './chatPage-reducer'
import { ChatMessage } from '../../../dal/api_chat/ChatService'

function createSocketConnections(producerId: number, consumerId: number) {
  return io('http://localhost:5001', { query: { producerId, consumerId } })
}

const createEventProviderChannel = (eventProvider: Socket) => {
  return eventChannel((emit) => {
    const receivedMessageHandler = (message: ChatMessage) => {
      emit(message)
    }
    eventProvider.on('SERVER_SEND_MESSAGE', receivedMessageHandler)

    return () => {
      eventProvider.off('SERVER_SEND_MESSAGE', receivedMessageHandler)
    }
  })
}

export function* openChannelSagaWorker(action: ReturnType<typeof openChannelRequest>) {
  const { producerId, consumerId } = action.payload
  const eventProvider: Socket = yield call(createSocketConnections, producerId, consumerId)
  yield put(setSocketChannel({ socketChannel: eventProvider }))

  // @ts-ignore
  const eventProviderChannel = yield call(createEventProviderChannel, eventProvider)

  try {
    while (true) {
      const payload: ChatMessage = yield take(eventProviderChannel)
      yield put(addMessage({ message: payload }))
    }
  } catch (err) {
    if (err instanceof Error) {
      yield put(fetchError({ errorMsg: err.message }))
    }
  }
}

export const openChannelRequest = (producerId: number, consumerId: number) => ({
  type: 'CHAT_PAGE/OPEN_CHANNEL',
  payload: {
    producerId,
    consumerId,
  },
})
