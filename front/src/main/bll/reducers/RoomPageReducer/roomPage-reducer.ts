import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchRoomResponse, RoomType } from '../../../dal/api_client/API'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialStateRoomPage = {
  loadingStatus: LoadingStatuses.IDLE,
  errorMsg: '',
  rooms: [],
  pageSize: 5,
  totalRoomsCount: 0,
  currentPage: 1,
}

const roomPageSlice = createSlice({
  name: 'roomPage',
  initialState,
  reducers: {
    fetchCurrentRoom(state) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.errorMsg = ''
    },
    fetchCurrentRoomSuccess(state, action: PayloadAction<{ rooms: FetchRoomResponse }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.rooms = action.payload.rooms
    },
    fetchCurrentRoomError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    setTotalRoomsCount(state, action: PayloadAction<{ totalRoomsCount: number }>) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.totalRoomsCount = action.payload.totalRoomsCount
    },
    setCurrentPage(state, action: PayloadAction<{ currentPage: number }>) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.currentPage = action.payload.currentPage
    },
  },
})
export const { fetchCurrentRoom, fetchCurrentRoomSuccess, fetchCurrentRoomError, setTotalRoomsCount, setCurrentPage } =
  roomPageSlice.actions
export const RoomPageReducer = roomPageSlice.reducer

//types

export type InitialStateRoomPage = {
  loadingStatus: RoomPageLoadingStatus
  errorMsg: string
  rooms: RoomType[]
  pageSize: number
  totalRoomsCount: number
  currentPage: number
}

export type RoomPageLoadingStatus =
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR
