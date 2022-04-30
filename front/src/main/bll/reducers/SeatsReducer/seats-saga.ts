import { call, put } from 'redux-saga/effects'
import moment from 'moment'
import {
  FetchRoomsResponse,
  FetchSeatsResponse,
  SeatResponse,
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
// eslint-disable-next-line no-unused-vars
// export const dayMaper = (arr: any[]) => {
//   return arr.map((el: any, index: number) => {
//     // eslint-disable-next-line no-unused-vars
//     return {
//       date: moment().add(index, 'days').format(),
//       ...el,
//     }
//   })
// }

export const dayMapper = (arr: SeatResponse[]) => {
  return arr.reduce((acc: SeatResponse, el, index: number) => {
    const date = moment().add(index, 'days').format('MM DD YYYY')
    console.log(acc)
    return {
      ...acc,
      [date]: el,
    }
  }, {} as SeatResponse)
}

export function* seatsSearchSagaWorker(action: ReturnType<typeof searchSeatsRequest>) {
  try {
    yield put(seatsStart())
    const searchSeatsResponse: AxiosResponse<SeatsSearchResponse> = yield call(
      SeatsAPI.fetchSeatsFree,
      action.newSeatsSearch
    )
    // @ts-ignore
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
