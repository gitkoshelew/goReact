import { call, put } from 'redux-saga/effects'
import {
  fetchError,
  fetchInitMessagesSuccess,
  fetchStart,
  fetchUsersSuccess,
  setCurrentConversation,
} from './chatPage-reducer'
import {
  ChatAPI,
  CreateConversationResponse,
  FetchInitMessagesResponse,
  FetchUsersResponse,
  GetConversationResponse,
} from '../../../dal/api_chat/ChatService'
import { AxiosResponse } from 'axios'

export function* fetchUsersSagaWorker() {
  try {
    yield put(fetchStart)
    const { data }: AxiosResponse<FetchUsersResponse> = yield call(ChatAPI.fetchUsers)
    yield put(fetchUsersSuccess({ users: data }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(fetchError({ errorMsg: err.message }))
    }
  }
}

export const fetchUsersRequest = () => ({
  type: 'CHAT_PAGE/FETCH_USERS_SAGA',
})

export function* fetchInitMessagesSagaWorker(action: FetchInitMessagesRequestType) {
  try {
    yield put(fetchStart())
    const { data }: AxiosResponse<FetchInitMessagesResponse> = yield call(
      ChatAPI.fetchInitMessages,
      action.conversationId
    )
    yield put(fetchInitMessagesSuccess({ initMessages: data }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(fetchError({ errorMsg: err.message }))
    }
  }
}

export const fetchInitMessagesRequest = (conversationId: number) => ({
  type: 'CHAT_PAGE/FETCH_INIT_MESSAGES_SAGA',
  conversationId,
})

type FetchInitMessagesRequestType = ReturnType<typeof fetchInitMessagesRequest>

export function* getConversationSagaWorker(action: GetConversationRequestType) {
  try {
    yield put(fetchStart())
    const foundedConversationResponse: AxiosResponse<GetConversationResponse> = yield call(
      ChatAPI.getConversation,
      action.producerId,
      action.consumerId
    )
    if (!foundedConversationResponse.data) {
      const createdConversationResponse: AxiosResponse<CreateConversationResponse> = yield call(
        ChatAPI.createConversation,
        action.producerId,
        action.consumerId
      )
      yield put(setCurrentConversation({ conversationId: createdConversationResponse.data.id }))
      yield put(fetchInitMessagesSuccess({ initMessages: [] }))
    } else {
      yield put(setCurrentConversation({ conversationId: foundedConversationResponse.data.id }))
      const initMessagesForConversationResponse: AxiosResponse<FetchInitMessagesResponse> = yield call(
        ChatAPI.fetchInitMessages,
        foundedConversationResponse.data.id
      )
      const initMessagesForConversation = initMessagesForConversationResponse.data
      yield put(fetchInitMessagesSuccess({ initMessages: initMessagesForConversation }))
    }
  } catch (err) {
    if (err instanceof Error) {
      yield put(fetchError({ errorMsg: err.message }))
    }
  }
}

export const getConversationRequest = (producerId: number, consumerId: number) => ({
  type: 'CHAT_PAGE/GET_CONVERSATION_SAGA',
  producerId,
  consumerId,
})

type GetConversationRequestType = ReturnType<typeof getConversationRequest>
