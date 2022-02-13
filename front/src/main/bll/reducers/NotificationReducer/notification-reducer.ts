import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Socket } from 'socket.io-client'

const initialState: InitialStateNotification = {
  socketChannel: null,
  isOpened: false,
  allNotifications: [],
  currentNotification: null,
}

const notificationSlice = createSlice({
  name: 'notification',
  initialState,
  reducers: {
    toggleNotification(state, action: PayloadAction<{ isOpen: boolean }>) {
      state.isOpened = action.payload.isOpen
    },
    addNotificationToQueue(state, action: PayloadAction<NotificationData>) {
      state.allNotifications.push(action.payload)
    },
    removeNotificationFromQueue(state) {
      state.allNotifications.shift()
    },
    setCurrentNotification(state, action: PayloadAction<NotificationData>) {
      state.currentNotification = action.payload
    },
    setNotificationSocketChannel(state, action: PayloadAction<{ socketChannel: Socket | null }>) {
      // @ts-ignore
      state.socketChannel = action.payload.socketChannel
    },
  },
})

export const NotificationReducer = notificationSlice.reducer
export const {
  toggleNotification,
  addNotificationToQueue,
  removeNotificationFromQueue,
  setCurrentNotification,
  setNotificationSocketChannel,
} = notificationSlice.actions

//types

type InitialStateNotification = {
  socketChannel: Socket | null
  isOpened: boolean
  allNotifications: NotificationData[]
  currentNotification: NotificationData | null
}

export type NotificationData = {
  toUser: number | null
  type: NotificationType
  reason: string
  description: string
}

type NotificationType = 'error' | 'warning' | 'info' | 'success'
