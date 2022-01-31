import { call, put } from 'redux-saga/effects'
import {
    fetchCurrentRoom, fetchCurrentRoomError,
    fetchCurrentRoomSuccess
} from './roomPage-reducer'

import { AxiosResponse } from 'axios'

import { FetchRoomResponse, RoomPageAPI } from "../../../dal/api_client/API";

export function* RoomPageSagaWorker() {
    try {
        yield put(fetchCurrentRoom)
        const { data }: AxiosResponse<FetchRoomResponse> = yield call(RoomPageAPI.getAllRooms)
        yield put(fetchCurrentRoomSuccess({ rooms: data }))
    } catch (err) {
        if (err instanceof Error) {
            yield put(fetchCurrentRoomError({ errorMsg: err.message }))
        } else {
            yield put(fetchCurrentRoomError({ errorMsg: 'asdRoomError' }))
        }
    }
}

export const fetchRoomRequest = () => ({
    type: 'ROOM_PAGE/FETCH_ROOM_SAGA',
})
