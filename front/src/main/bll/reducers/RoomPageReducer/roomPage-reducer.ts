import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchRoomResponse, RoomType } from "../../../dal/api_client/API";

enum RoomPageLoadingStatuses {
    IDLE = 'IDLE',
    LOADING = 'LOADING',
    SUCCESS = 'SUCCESS',
    ERROR = 'ERROR',
}

const initialState: InitialStateRoomPage = {
    loadingStatus: RoomPageLoadingStatuses.IDLE,
    errorMsg: '',
    rooms: [],
    pageSize: 5,
    totalRoomsCount: 0,
    currentPage: 1,
}

const roomPageSlice = createSlice({
    name: 'roomPage',
    initialState,
    reducers: {
        fetchCurrentRoom(state) {
            state.loadingStatus = RoomPageLoadingStatuses.LOADING
            state.errorMsg = ''
        },
        fetchCurrentRoomSuccess(state, action: PayloadAction<{ rooms: FetchRoomResponse }>) {
            state.loadingStatus = RoomPageLoadingStatuses.SUCCESS
            state.rooms = action.payload.rooms
        },
        fetchCurrentRoomError(state, action: PayloadAction<{ errorMsg: string }>) {
            state.loadingStatus = RoomPageLoadingStatuses.ERROR
            state.errorMsg = action.payload.errorMsg
        },
        setTotalRoomsCount(state, action: PayloadAction<{ totalRoomsCount: number }>) {
            state.loadingStatus = RoomPageLoadingStatuses.LOADING
            state.totalRoomsCount = action.payload.totalRoomsCount
        },
        setCurrentPage(state, action: PayloadAction<{ currentPage: number }>) {
            state.loadingStatus = RoomPageLoadingStatuses.LOADING
            state.currentPage = action.payload.currentPage
        },
    },
})
export const {
    fetchCurrentRoom,
    fetchCurrentRoomSuccess,
    fetchCurrentRoomError,
    setTotalRoomsCount,
    setCurrentPage
} = roomPageSlice.actions
export const RoomPageReducer = roomPageSlice.reducer

//types

export type InitialStateRoomPage = {
    loadingStatus: RoomPageLoadingStatus
    errorMsg: string,
    rooms: RoomType[],
    pageSize: number,
    totalRoomsCount: number,
    currentPage: number,
}

export type RoomPageLoadingStatus =
    | RoomPageLoadingStatuses.IDLE
    | RoomPageLoadingStatuses.LOADING
    | RoomPageLoadingStatuses.SUCCESS
    | RoomPageLoadingStatuses.ERROR
