import {createSlice} from '@reduxjs/toolkit'


const hotelsSlice = createSlice({
    name: 'hotelPage',
    initialState: {
        hotels: [],
        isLoading: false
    },
    reducers: {
        getHotelsFetch: (state) => {
            state.isLoading = true;
        },
        getHotelsSuccess:(state, action)=>{
           state.hotels = action.payload
           state.isLoading = false;
        },
        getCatsFailure:(state)=>{
            state.isLoading=false;
        }

    }
})

export const hotelsReducer = hotelsSlice.reducer


export const {getHotelsFetch,getHotelsSuccess,getCatsFailure}=hotelsSlice.actions;


//types
export type HotelType = {
    id: string
    name: string
}


export type InitialHotelsPageStateType = {
    hotels: HotelType[]
}
type AllHotelsPageActionType = any
