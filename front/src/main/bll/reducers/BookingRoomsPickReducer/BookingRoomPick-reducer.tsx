import { createSlice, PayloadAction } from '@reduxjs/toolkit'

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
    setLoadingStatus(state, action: PayloadAction<{ newStatus: LoadingStatusBookingPickType }>) {
      state.loadingStatus = action.payload.newStatus
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

export const { changeActualDay, setIsRent, addOrderedRoom, changeRoomStatus, deleteOrderedRoom, setLoadingStatus } =
  BookingRoomPickSlice.actions

//types

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }
export type OrderedRoomsType = { id: string; orderedRoomType: RentRoomType }

type InitialStateType = {
  isRent: IsRentType[]
  actualDay: string | Date
  orderedRoomBasket: OrderedRoomsType[]
  loadingStatus: LoadingStatusBookingPickType
}

export type LoadingStatusBookingPickType = 'loading' | 'success' | 'error'
export type RentRoomType = 'firstRoom' | 'secondRoom'
