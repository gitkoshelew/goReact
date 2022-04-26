import { createSlice } from '@reduxjs/toolkit'

enum SocialNetworkAuthLoadingStatuses {
  IDLE = 'IDLE',
  LOADING = 'LOADING',
  SUCCESS = 'SUCCESS',
  ERROR = 'ERROR',
}

const initialState: InitialStateSeats = {
  loadingStatus: SocialNetworkAuthLoadingStatuses.IDLE,
}
const socialNetworkAuthSlice = createSlice({
  name: 'socialNetworkAuth',
  initialState,
  reducers: {
    socialNetworkAuthStart(state) {
      state.loadingStatus = SocialNetworkAuthLoadingStatuses.LOADING
    },
    socialNetworkAuthSuccess(state) {
      state.loadingStatus = SocialNetworkAuthLoadingStatuses.SUCCESS
    },
  },
})

export const SocialNetworkAuthReducer = socialNetworkAuthSlice.reducer
export const { socialNetworkAuthStart, socialNetworkAuthSuccess } = socialNetworkAuthSlice.actions

//types

type InitialStateSeats = {
  loadingStatus: BookingPaymentLoadingStatus
}

export type BookingPaymentLoadingStatus =
  | SocialNetworkAuthLoadingStatuses.IDLE
  | SocialNetworkAuthLoadingStatuses.LOADING
  | SocialNetworkAuthLoadingStatuses.SUCCESS
  | SocialNetworkAuthLoadingStatuses.ERROR
