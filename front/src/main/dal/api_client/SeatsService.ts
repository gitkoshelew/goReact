import { AxiosResponse } from 'axios'
import { $api } from './API'

type Pet = 'Cat' | 'Dog'

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

export type FetchSeatsResponse = Seat[]

export const SeatsAPI = {
  async fetchRooms(): Promise<AxiosResponse<FetchRoomsResponse>> {
    return $api.get(`api/rooms`)
  },
  async fetchSeats(): Promise<AxiosResponse<FetchSeatsResponse>> {
    return $api.get(`api/seats`)
  },
}
