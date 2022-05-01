import { AxiosResponse } from 'axios'
import { $api } from './API'

export type Pet = 'cat' | 'dog'

export type Room = {
  roomId: number
  roomNum: number
  petType: Pet
  hotelId: number
  roomPhotoUrl: string
}

export type FetchRoomsResponse = {
  rooms: Room[]
  totalCount: number
}

export type Seat = {
  seatId: number
  roomId: number
  description: string
  rentFrom: Date
  rentTo: Date
}

export type SeatResponse = {
  day: number
  seatIds: number[]
  totalCount?: number
}

export type SeatSearch = {
  hotelId: number
  petType: Pet | string
}

export type SeatResponseWithNewKey = {
  [key: string]: { day: number; seatIds: number[]; totalCount: number }
}

export type FetchSeatsResponse = Seat[]
export type SeatsSearchResponse = SeatResponse[]

export const SeatsAPI = {
  async fetchRooms(): Promise<AxiosResponse<FetchRoomsResponse>> {
    return $api.get(`api/rooms`)
  },
  async fetchSeats(): Promise<AxiosResponse<FetchSeatsResponse>> {
    return $api.get(`api/seats`)
  },
  async fetchSeatsFree(newSeatsSearch: SeatSearch): Promise<AxiosResponse<SeatsSearchResponse>> {
    const res = await $api.post(`api/seats/search/free`, newSeatsSearch)
    return res
  },
}
