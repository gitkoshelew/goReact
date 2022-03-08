import { call, put } from 'redux-saga/effects'
import { FetchRoomsResponse, FetchSeatsResponse, SeatsAPI } from '../../../dal/api_client/SeatsService'
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
