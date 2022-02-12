import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Socket } from 'socket.io-client'

const initialState: InitialStateNotification = {
  socketChannel: null,
  isOpened: false,
  data: {
    toUser: null,
    type: 'info',
    reason: '',
    description: '',
  },
}

const notificationSlice = createSlice({
  name: 'notification',
  initialState,
  reducers: {
    toggleNotification(state, action: PayloadAction<{ isOpen: boolean }>) {
      state.isOpened = action.payload.isOpen
    },
    setNotificationData(state, action: PayloadAction<NotificationData>) {
      state.data = action.payload
    },
    setNotificationSocketChannel(state, action: PayloadAction<{ socketChannel: Socket | null }>) {
      // @ts-ignore
      state.socketChannel = action.payload.socketChannel
    },
  },
})

export const NotificationReducer = notificationSlice.reducer
export const { toggleNotification, setNotificationData, setNotificationSocketChannel } = notificationSlice.actions

//types

type InitialStateNotification = {
  socketChannel: Socket | null
  isOpened: boolean
  data: NotificationData
}

export type NotificationData = {
  toUser: number | null
  type: NotificationType
  reason: string
  description: string
}

type NotificationType = 'error' | 'warning' | 'info' | 'success'
