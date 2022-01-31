import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchRoomResponse, RoomType } from "../../../dal/api_client/API";


const initialState: InitialStateRoomPage = {
    loadingStatus: 'idle',
    errorMsg: '',
    rooms: []

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

    },
})
export const {
    fetchCurrentRoom,
    fetchCurrentRoomSuccess,
    fetchCurrentRoomError
} = roomPageSlice.actions
export const RoomPageReducer = roomPageSlice.reducer

//types

export type InitialStateRoomPage = {
    loadingStatus: RoomPageLoadingStatus
    errorMsg: string,
    rooms: RoomType[],
}

export type RoomPageLoadingStatus = 'idle' | 'loading' | 'success' | 'error'
