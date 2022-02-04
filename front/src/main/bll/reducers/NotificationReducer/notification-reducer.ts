import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: InitialStateNotification = {
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
  },
})

export const NotificationReducer = notificationSlice.reducer
export const { toggleNotification, setNotificationData } = notificationSlice.actions

//types

type InitialStateNotification = {
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
