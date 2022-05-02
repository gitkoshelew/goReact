import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Room, Seat, SeatResponseWithNewKey } from '../../../dal/api_client/SeatsService'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialStateSeats = {
  loadingStatus: LoadingStatuses.IDLE,
  rooms: [],
  seats: [],
  seatsSearch: {},
  successMsg: '',
  errorMsg: '',
}

const seatsSlice = createSlice({
  name: 'seats',
  initialState,
  reducers: {
    seatsStart(state) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.successMsg = ''
      state.errorMsg = ''
    },
    setRooms(state, action: PayloadAction<{ rooms: Room[] }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.rooms = action.payload.rooms
    },
    setSeats(state, action: PayloadAction<{ seats: Seat[] }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.seats = action.payload.seats
    },
    setSeatsSearch(state, action: PayloadAction<{ seatsSearch: SeatResponseWithNewKey }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.seatsSearch = action.payload.seatsSearch
    },

    seatsError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    seatsMessage(state, action: PayloadAction<{ successMsg: string }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.successMsg = action.payload.successMsg
    },
  },
})

export const SeatsReducer = seatsSlice.reducer
export const { seatsStart, setRooms, setSeats, setSeatsSearch, seatsMessage, seatsError } = seatsSlice.actions

//types

type InitialStateSeats = {
  loadingStatus: SeatsLoadingStatus
  rooms: Room[]
  seats: Seat[]
  seatsSearch: SeatResponseWithNewKey
  successMsg: string
  errorMsg: string
}

export type SeatsLoadingStatus =
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR
