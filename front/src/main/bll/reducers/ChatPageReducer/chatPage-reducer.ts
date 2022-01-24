import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchUsersResponse, ChatUser, ChatMessage } from '../../../dal/api_chat/ChatService'

const initialState: InitialStateChatPage = {
  loadingStatus: 'idle',
  errorMsg: '',
  users: [],
  messages: [],
  conversationId: null,
}

const chatPageSlice = createSlice({
  name: 'chatPage',
  initialState,
  reducers: {
    fetchStart(state) {
      state.loadingStatus = 'loading'
      state.errorMsg = ''
    },
    fetchError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = 'error'
      state.errorMsg = action.payload.errorMsg
    },
    fetchUsersSuccess(state, action: PayloadAction<{ users: FetchUsersResponse }>) {
      state.loadingStatus = 'success'
      state.users = action.payload.users
    },
    fetchInitMessagesSuccess(state, action: PayloadAction<{ initMessages: ChatMessage[] }>) {
      state.messages = action.payload.initMessages
    },
    setCurrentConversation(state, action: PayloadAction<{ conversationId: number }>) {
      state.loadingStatus = 'success'
      state.conversationId = action.payload.conversationId
    },
  },
})

export const ChatPageReducer = chatPageSlice.reducer
export const { fetchStart, fetchUsersSuccess, fetchError, fetchInitMessagesSuccess, setCurrentConversation } =
  chatPageSlice.actions

//types

export type InitialStateChatPage = {
  loadingStatus: ChatPageLoadingStatus
  errorMsg: string
  users: ChatUser[]
  messages: ChatMessage[]
  conversationId: number | null
}

export type ChatPageLoadingStatus = 'idle' | 'loading' | 'success' | 'error'
