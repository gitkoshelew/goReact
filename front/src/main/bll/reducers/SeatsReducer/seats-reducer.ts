import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Room, Seat, SeatSearch } from '../../../dal/api_client/SeatsService'

const initialState: InitialStateSeats = {
  rooms: [],
  seats: [],
  seatsSearch: [],
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
    setSeatsSearch(state, action: PayloadAction<{ seatsSearch: SeatSearch[] }>) {
      state.seatsSearch = action.payload.seatsSearch
    },
  },
})

export const SeatsReducer = seatsSlice.reducer
export const { setRooms, setSeats, setSeatsSearch } = seatsSlice.actions

//types

type InitialStateSeats = {
  rooms: Room[]
  seats: Seat[]
  seatsSearch: SeatSearch[]
}
