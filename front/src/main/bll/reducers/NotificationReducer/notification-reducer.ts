import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Socket } from 'socket.io-client'
import { NotificationRaw } from './socketChannel'

const initialState: InitialStateNotification = {
  socketChannel: null,
  notifications: [],
}

const notificationSlice = createSlice({
  name: 'notification',
  initialState,
  reducers: {
    addNotification(state, action: PayloadAction<NotificationRaw>) {
      state.notifications.push(action.payload)
    },
    removeNotification(state, action: PayloadAction<string>) {
      state.notifications = state.notifications.filter((notification) => notification.content.id !== action.payload)
    },
    setNotificationSocketChannel(state, action: PayloadAction<{ socketChannel: Socket | null }>) {
      // @ts-ignore
      state.socketChannel = action.payload.socketChannel
    },
  },
})

export const NotificationReducer = notificationSlice.reducer
export const { addNotification, removeNotification, setNotificationSocketChannel } = notificationSlice.actions

//types

type InitialStateNotification = {
  socketChannel: Socket | null
  notifications: NotificationRaw[]
}

export type NotificationData = {
  id: string
  toUser: number | null
  type: NotificationType
  reason: string
  description: string
}

type NotificationType = 'error' | 'warning' | 'info' | 'success'
