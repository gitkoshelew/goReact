import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: InitialStateType = {
  isRent: [
    { id: '120121', firstRoom: true, secondRoom: false },
    { id: '120121', firstRoom: true, secondRoom: false },
    { id: '120221', firstRoom: true, secondRoom: false },
    { id: '120321', firstRoom: true, secondRoom: false },
    { id: '120421', firstRoom: true, secondRoom: false },
    { id: '120521', firstRoom: true, secondRoom: false },
    { id: '120621', firstRoom: true, secondRoom: false },
    { id: '120721', firstRoom: true, secondRoom: false },
    { id: '120821', firstRoom: true, secondRoom: false },
    { id: '120921', firstRoom: true, secondRoom: false },
    { id: '121021', firstRoom: true, secondRoom: false },
    { id: '121121', firstRoom: true, secondRoom: false },
    { id: '121221', firstRoom: true, secondRoom: false },
    { id: '121321', firstRoom: true, secondRoom: false },
    { id: '121421', firstRoom: true, secondRoom: false },
    { id: '121521', firstRoom: true, secondRoom: false },
    { id: '121621', firstRoom: true, secondRoom: false },
    { id: '121721', firstRoom: true, secondRoom: false },
    { id: '121821', firstRoom: false, secondRoom: false },
    { id: '121921', firstRoom: true, secondRoom: false },
    { id: '122021', firstRoom: false, secondRoom: true },
    { id: '122121', firstRoom: false, secondRoom: false },
    { id: '122221', firstRoom: true, secondRoom: true },
    { id: '122321', firstRoom: false, secondRoom: true },
    { id: '122421', firstRoom: true, secondRoom: false },
    { id: '122521', firstRoom: true, secondRoom: false },
    { id: '122621', firstRoom: true, secondRoom: true },
    { id: '122721', firstRoom: false, secondRoom: true },
    { id: '122821', firstRoom: false, secondRoom: true },
    { id: '122921', firstRoom: true, secondRoom: false },
    { id: '123021', firstRoom: false, secondRoom: true },
    { id: '123121', firstRoom: true, secondRoom: false },
  ],
  actualDay: new Date() as Date | string,
  orderedRoomBasket: [],
}

const BookingRoomPickSlice = createSlice({
  name: 'BookingRoomPick',
  initialState,
  reducers: {
    changeActualDay(state, action: PayloadAction<{ newActualDay: string }>) {
      state.actualDay = action.payload.newActualDay
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

export const { changeActualDay, addNewIsRent, addOrderedRoom, changeRoomStatus, deleteOrderedRoom } =
  BookingRoomPickSlice.actions

//types

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }
export type OrderedRoomsType = { id: string; orderedRoomType: RentRoomType }

type InitialStateType = {
  isRent: IsRentType[]
  actualDay: string | Date
  orderedRoomBasket: OrderedRoomsType[]
}

export type RentRoomType = 'firstRoom' | 'secondRoom'
