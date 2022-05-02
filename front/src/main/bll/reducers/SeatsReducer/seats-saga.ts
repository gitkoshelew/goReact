import { call, put } from 'redux-saga/effects'
import moment from 'moment'
import {
  FetchRoomsResponse,
  FetchSeatsResponse,
  SeatResponse,
  SeatResponseWithNewKey,
  SeatsAPI,
  SeatSearch,
  SeatsSearchResponse,
} from '../../../dal/api_client/SeatsService'
import { seatsError, seatsMessage, seatsStart, setRooms, setSeats, setSeatsSearch } from './seats-reducer'
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

export const dayMapper = (arr: SeatResponse[]): SeatResponseWithNewKey => {
  return arr.reduce((acc: {}, el: SeatResponse) => {
    const date = moment().add(el.day, 'days').format('MM DD YYYY')
    return {
      ...acc,
      [date]: el,
    }
  }, {})
}

export function* seatsSearchSagaWorker(action: ReturnType<typeof searchSeatsRequest>) {
  try {
    yield put(seatsStart())
    const searchSeatsResponse: AxiosResponse<SeatsSearchResponse> = yield call(
      SeatsAPI.fetchSeatsFree,
      action.newSeatsSearch
    )
    yield put(setSeatsSearch({ seatsSearch: dayMapper(searchSeatsResponse.data) }))
    yield put(seatsMessage({ successMsg: searchSeatsResponse.statusText }))
  } catch (err) {
    if (err) {
      yield put(seatsError({ errorMsg: 'Looks like something went wrong, please try again later' }))
    } else {
      yield put(seatsError({ errorMsg: 'Looks like something went wrong, please try again later' }))
    }
  }
}

export const searchSeatsRequest = (newSeatsSearch: SeatSearch) => ({
  type: 'SEATS/SEARCH_SEATS_SAGA',
  newSeatsSearch,
})
