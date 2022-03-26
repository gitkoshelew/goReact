import { call, put } from 'redux-saga/effects'
import {
  FetchRoomsResponse,
  FetchSeatsResponse,
  SeatsAPI,
  SeatSearch,
  SeatsSearchResponse,
} from '../../../dal/api_client/SeatsService'
import { setRooms, setSeats } from './seats-reducer'
import { AxiosResponse } from 'axios'

export function* fetchSeatsSagaWorker() {
  const roomsResponse: AxiosResponse<FetchRoomsResponse> = yield call(SeatsAPI.fetchRooms)
  yield put(setRooms({ rooms: roomsResponse.data.rooms }))
  const seatsResponse: AxiosResponse<FetchSeatsResponse> = yield call(SeatsAPI.fetchSeats)
  yield put(setSeats({ seats: seatsResponse.data }))
}

export const fetchSeatsRequest = () => ({
  type: 'SEATS/FETCH_SETS',
})

export function* seatsSearchSagaWorker(action: ReturnType<typeof searchSeatsRequest>) {
  const searchSeatsResponse: AxiosResponse<SeatsSearchResponse> = yield call(
    SeatsAPI.fetchSeatsFree,
    action.newSeatsSearch
  )
}

export const searchSeatsRequest = (newSeatsSearch: SeatSearch) => ({
  type: 'SEATS/SEARCH_SEATS_SAGA',
  newSeatsSearch,
})
