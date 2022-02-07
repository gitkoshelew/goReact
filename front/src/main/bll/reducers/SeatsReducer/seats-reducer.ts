import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Room, Seat } from '../../../dal/api_client/SeatsService'

const initialState: InitialStateSeats = {
  rooms: [],
  seats: [],
}

const seatsSlice = createSlice({
  name: 'seats',
  initialState,
  reducers: {
    setRooms(state, action: PayloadAction<{ rooms: Room[] }>) {
      state.rooms = action.payload.rooms
    },
    setSeats(state, action: PayloadAction<{ seats: Seat[] }>) {
      state.seats = action.payload.seats
    },
  },
})

export const SeatsReducer = seatsSlice.reducer
export const { setRooms, setSeats } = seatsSlice.actions

//types

type InitialStateSeats = {
  rooms: Room[]
  seats: Seat[]
}
