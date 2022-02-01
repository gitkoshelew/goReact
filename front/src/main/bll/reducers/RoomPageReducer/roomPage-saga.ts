import { call, put } from 'redux-saga/effects'
import {
    fetchCurrentRoom, fetchCurrentRoomError,
    fetchCurrentRoomSuccess, setCurrentPage, setTotalRoomsCount
} from './roomPage-reducer'

import { AxiosResponse } from 'axios'

import { FetchRoomResponse, RoomPageAPI } from "../../../dal/api_client/API";

export function* RoomPageSagaWorker(action: RoomRequest) {
    try {
        yield put(setCurrentPage({ currentPage: action.currentPage }))
        yield put(fetchCurrentRoom)
        const { data }: AxiosResponse<FetchRoomResponse> = yield call(RoomPageAPI.getRooms, action.currentPage, action.pageSize)
        yield put(fetchCurrentRoomSuccess({ rooms: data }))
        yield put(setTotalRoomsCount)
    } catch (err) {
        if (err instanceof Error) {
            yield put(fetchCurrentRoomError({ errorMsg: err.message }))
        } else {
            yield put(fetchCurrentRoomError({ errorMsg: 'asdRoomError' }))
        }
    }
}

export const fetchRoomRequest = (currentPage: number, pageSize: number) => ({
    type: 'ROOM_PAGE/FETCH_ROOM_SAGA',
    currentPage,
    pageSize
})

type RoomRequest = ReturnType<typeof fetchRoomRequest>
