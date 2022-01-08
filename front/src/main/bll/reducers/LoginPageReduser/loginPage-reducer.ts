import {createSlice, PayloadAction} from '@reduxjs/toolkit';


const initialState: InitialStateLoginPageType = {
    loadingStatus: 'onWaiting',
    userForLogin: null,
}


const loginPageSlice = createSlice({
    name: 'loginPage',
    initialState,
    reducers: {
        changeUserParams(state, action: PayloadAction<{ user: UserForLoginType }>) {
            state.userForLogin = action.payload.user
        }
    }
})

export const LoginPageReducer = loginPageSlice.reducer
export const {changeUserParams} = loginPageSlice.actions




//types

export type InitialStateLoginPageType = {
    loadingStatus: LoadingStatusType,
    userForLogin: UserForLoginType
}


export type LoadingStatusType = 'onWaiting' | 'loading' | 'success' | 'error'
export type UserForLoginType = { email: string, password: string } | null