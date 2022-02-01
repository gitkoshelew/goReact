import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchRoomResponse, RoomType } from "../../../dal/api_client/API";


const initialState: InitialStateRoomPage = {
    loadingStatus: 'idle',
    errorMsg: '',
    rooms: [],
    pageSize: 5,
    totalRoomsCount: 0,
    currentPage: 2,
}

const roomPageSlice = createSlice({
    name: 'roomPage',
    initialState,
    reducers: {
        fetchCurrentRoom(state) {
            state.loadingStatus = 'loading'
            state.errorMsg = ''
        },
        fetchCurrentRoomSuccess(state, action: PayloadAction<{ rooms: FetchRoomResponse }>) {
            state.loadingStatus = 'success'
            state.rooms = action.payload.rooms
        },
        fetchCurrentRoomError(state, action: PayloadAction<{ errorMsg: string }>) {
            state.loadingStatus = 'error'
            state.errorMsg = action.payload.errorMsg
        },
        setTotalRoomsCount(state, action: PayloadAction<{ totalRoomsCount: number }>) {
            state.loadingStatus = 'loading'
            state.totalRoomsCount = action.payload.totalRoomsCount
        },
        setCurrentPage(state, action: PayloadAction<{ currentPage: number }>) {
            state.loadingStatus = 'loading'
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

export type RoomPageLoadingStatus = 'idle' | 'loading' | 'success' | 'error'
