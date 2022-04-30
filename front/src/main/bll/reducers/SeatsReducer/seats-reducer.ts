import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Room, Seat, SeatResponse } from '../../../dal/api_client/SeatsService'

enum SeatsLoadingStatuses {
  IDLE = 'IDLE',
  LOADING = 'LOADING',
  SUCCESS = 'SUCCESS',
  ERROR = 'ERROR',
}

const initialState: InitialStateSeats = {
  loadingStatus: SeatsLoadingStatuses.IDLE,
  rooms: [],
  seats: [],
  seatsSearch: [],
  successMsg: '',
  errorMsg: '',
}

const seatsSlice = createSlice({
  name: 'seats',
  initialState,
  reducers: {
    seatsStart(state) {
      state.loadingStatus = SeatsLoadingStatuses.LOADING
      state.successMsg = ''
      state.errorMsg = ''
    },
    setRooms(state, action: PayloadAction<{ rooms: Room[] }>) {
      state.loadingStatus = SeatsLoadingStatuses.SUCCESS
      state.rooms = action.payload.rooms
    },
    setSeats(state, action: PayloadAction<{ seats: Seat[] }>) {
      state.loadingStatus = SeatsLoadingStatuses.SUCCESS
      state.seats = action.payload.seats
    },
    setSeatsSearch(state, action: PayloadAction<{ seatsSearch: SeatResponse[] }>) {
      state.loadingStatus = SeatsLoadingStatuses.SUCCESS
      state.seatsSearch = action.payload.seatsSearch
    },
    seatsError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = SeatsLoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    seatsMessage(state, action: PayloadAction<{ successMsg: string }>) {
      state.loadingStatus = SeatsLoadingStatuses.SUCCESS
      state.successMsg = action.payload.successMsg
    },
  },
})

export const SeatsReducer = seatsSlice.reducer
export const { seatsStart, setRooms, setSeats, setSeatsSearch, seatsMessage, seatsError } =
  seatsSlice.actions

//types

type InitialStateSeats = {
  loadingStatus: SeatsLoadingStatus
  rooms: Room[]
  seats: Seat[]
  seatsSearch: SeatResponse[]
  successMsg: string
  errorMsg: string
}

export type SeatsLoadingStatus =
  | SeatsLoadingStatuses.IDLE
  | SeatsLoadingStatuses.LOADING
  | SeatsLoadingStatuses.SUCCESS
  | SeatsLoadingStatuses.ERROR
