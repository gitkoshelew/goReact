import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchUsersResponse, ChatUser, ChatMessage } from '../../../dal/api_chat/ChatService'
import { Socket } from 'socket.io-client'

enum ChatPageLoadingStatuses {
  IDLE = 'IDLE',
  LOADING = 'LOADING',
  SUCCESS = 'SUCCESS',
  ERROR = 'ERROR',
}

const initialState: InitialStateChatPage = {
  loadingStatus: ChatPageLoadingStatuses.IDLE,
  errorMsg: '',
  users: [],
  messages: [],
  conversationId: null,
  socketChannel: null,
  isConversationOpened: false,
}

const chatPageSlice = createSlice({
  name: 'chatPage',
  initialState,
  reducers: {
    fetchStart(state) {
      state.loadingStatus = ChatPageLoadingStatuses.LOADING
      state.errorMsg = ''
    },
    fetchError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = ChatPageLoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    fetchUsersSuccess(state, action: PayloadAction<{ users: FetchUsersResponse }>) {
      state.loadingStatus = ChatPageLoadingStatuses.SUCCESS
      state.users = action.payload.users
    },
    fetchInitMessagesSuccess(state, action: PayloadAction<{ initMessages: ChatMessage[] }>) {
      state.messages = action.payload.initMessages
    },
    setCurrentConversation(state, action: PayloadAction<{ conversationId: number }>) {
      state.loadingStatus = ChatPageLoadingStatuses.SUCCESS
      state.conversationId = action.payload.conversationId
    },
    setSocketChannel(state, action: PayloadAction<{ socketChannel: Socket | null }>) {
      // @ts-ignore
      state.socketChannel = action.payload.socketChannel
    },
    addMessage(state, action: PayloadAction<{ message: ChatMessage }>) {
      state.messages.push(action.payload.message)
    },
    setConversationOpened(state, action: PayloadAction<{ isOpened: boolean }>) {
      state.isConversationOpened = action.payload.isOpened
    },
  },
})

export const ChatPageReducer = chatPageSlice.reducer
export const {
  fetchStart,
  fetchUsersSuccess,
  fetchError,
  fetchInitMessagesSuccess,
  setCurrentConversation,
  addMessage,
  setSocketChannel,
  setConversationOpened,
} = chatPageSlice.actions

//types

export type InitialStateChatPage = {
  loadingStatus: ChatPageLoadingStatus
  errorMsg: string
  users: ChatUser[]
  messages: ChatMessage[]
  conversationId: number | null
  socketChannel: Socket | null
  isConversationOpened: boolean
}

export type ChatPageLoadingStatus =
  | ChatPageLoadingStatuses.IDLE
  | ChatPageLoadingStatuses.LOADING
  | ChatPageLoadingStatuses.SUCCESS
  | ChatPageLoadingStatuses.ERROR
