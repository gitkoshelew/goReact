import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IsRentType } from '../../../dal/API'

const initialState: InitialStateType = {
  isRent: [],
  loadingStatus: 'success',
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
    setIsRent(state, action: PayloadAction<{ newIsRent: IsRentType[] }>) {
      state.isRent = action.payload.newIsRent
    },
    addNewIsRent(state, action: PayloadAction<{ newIsRent: IsRentType[] }>) {
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
      state.loadingStatus = 'loading'
    },
    reqCalendarDataSucceeded(state, action: PayloadAction<{ newCalendarData: IsRentType[] }>) {
      state.loadingStatus = 'success'
      debugger
      state.isRent = action.payload.newCalendarData
    },
    reqCalendarDataFailed(state) {
      state.loadingStatus = 'error'
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
  isRent: IsRentType[]
  actualDay: string | Date
  orderedRoomBasket: OrderedRoomsType[]
  loadingStatus: LoadingStatusBookingPickType
}

export type LoadingStatusBookingPickType = 'loading' | 'success' | 'error'
export type RentRoomType = 'firstRoom' | 'secondRoom'
