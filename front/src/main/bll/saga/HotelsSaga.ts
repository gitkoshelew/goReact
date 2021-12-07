import {call, put, takeEvery} from 'redux-saga/effects'
import {getHotelsFetch, getHotelsSuccess} from '../redusers/Hotels-reducer';
import {PetHotelsService} from '../../dal/API';
import {AxiosResponse} from 'axios';

const api = new PetHotelsService()


export function* HotelsSaga(){
    yield takeEvery('HOTELS_SAGA/GETS_CATS_FETCH',hotelsWorker)
}




export const fetchHotelsAC=()=>({type:'HOTELS_SAGA/GETS_CATS_FETCH'})



function* hotelsWorker(){
    yield put(getHotelsFetch())
    const res:AxiosResponse<any> = yield call(api.getAllHotels)
    yield put(getHotelsSuccess(res))
}