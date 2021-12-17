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
  userName: '',
  userEmail: '',
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
    changeUserParams(state, action: PayloadAction<{ params: string; newTextParams: string }>) {
      switch (action.payload.params) {
        case 'name':
          state.userName = action.payload.newTextParams
          break
        case 'email':
          state.userEmail = action.payload.newTextParams
          break
        default:
          return state
      }
    },
  },
})

export const BookingRoomPickReducer = BookingRoomPickSlice.reducer

export const { changeActualDay, addNewIsRent, changeUserParams } = BookingRoomPickSlice.actions

//types

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }

type InitialStateType = {
  isRent: IsRentType[]
  actualDay: string | Date
  userName: string
  userEmail: string
}
