import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IsRent } from '../../../dal/api_client/API'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialStateType = {
  isRent: [],
  loadingStatus: LoadingStatuses.SUCCESS,
  actualDay: new Date(),
  orderedRoomBasket: [],
}

const BookingRoomPickSlice = createSlice({
  name: 'BookingRoomPick',
  initialState,
  reducers: {
    changeActualDay(state, action: PayloadAction<{ newActualDay: string }>) {
      state.actualDay = action.payload.newActualDay
    },
    setIsRent(state, action: PayloadAction<{ newIsRent: IsRent[] }>) {
      state.isRent = action.payload.newIsRent
    },
    addNewIsRent(state, action: PayloadAction<{ newIsRent: IsRent[] }>) {
      state.isRent = [...state.isRent, ...action.payload.newIsRent]
    },
    changeRoomStatus(state, action: PayloadAction<{ roomType: RentRoomType; dayId: string }>) {
      const isRentIndex = state.isRent.findIndex((t) => t.id === action.payload.dayId)
      const actualRoomType = action.payload.roomType
      actualRoomType === 'firstRoom'
        ? (state.isRent[isRentIndex].firstRoom = false)
        : (state.isRent[isRentIndex].secondRoom = false)
    },
    addOrderedRoom(state, action: PayloadAction<{ newOrderedRooms: OrderedRoomsType }>) {
      state.orderedRoomBasket.push(action.payload.newOrderedRooms)
    },
    reqCalendarDataStarted(state) {
      state.loadingStatus = LoadingStatuses.LOADING
    },
    reqCalendarDataSucceeded(state, action: PayloadAction<{ newCalendarData: IsRent[] }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.isRent = action.payload.newCalendarData
    },
    reqCalendarDataFailed(state) {
      state.loadingStatus = LoadingStatuses.ERROR
    },
    deleteOrderedRoom(state, action: PayloadAction<{ newOrderedRooms: OrderedRoomsType }>) {
      const roomForDeleteIndex = state.orderedRoomBasket.findIndex((t) => t.id === action.payload.newOrderedRooms.id)
      state.orderedRoomBasket.splice(roomForDeleteIndex, 1)
      const roomToChangeStatus = state.isRent.findIndex((t) => t.id === action.payload.newOrderedRooms.id)
      action.payload.newOrderedRooms.orderedRoomType === 'firstRoom'
        ? (state.isRent[roomToChangeStatus].firstRoom = true)
        : (state.isRent[roomToChangeStatus].secondRoom = true)
    },
  },
})

export const BookingRoomPickReducer = BookingRoomPickSlice.reducer

export const {
  changeActualDay,
  setIsRent,
  addOrderedRoom,
  changeRoomStatus,
  deleteOrderedRoom,
  reqCalendarDataStarted,
  reqCalendarDataSucceeded,
  reqCalendarDataFailed,
} = BookingRoomPickSlice.actions

//types

export type OrderedRoomsType = { id: string; orderedRoomType: RentRoomType }

type InitialStateType = {
  isRent: IsRent[]
  actualDay: string | Date
  orderedRoomBasket: OrderedRoomsType[]
  loadingStatus: LoadingStatusBookingPickType
}

export type LoadingStatusBookingPickType =
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR

export type RentRoomType = 'firstRoom' | 'secondRoom'
