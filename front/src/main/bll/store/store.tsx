import { combineReducers } from 'redux'
import { configureStore } from '@reduxjs/toolkit'
import createSagaMiddleware from 'redux-saga'
import { BookingRegFormReducer } from '../reducers/BookingRegFormReducer/BookingRegForm-reducer'
import { useDispatch } from 'react-redux'
import { takeEvery } from 'redux-saga/effects'
import { BookingUploadPetImgSagaWorker } from '../reducers/BookingRegFormReducer/BookindRegForm-saga'

const sagaMiddleware = createSagaMiddleware()

const rootReducer = combineReducers({
  BookingRegForm: BookingRegFormReducer,
})

export type RootReducerType = typeof rootReducer

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().prepend(sagaMiddleware),
})

export type AppRootStateType = ReturnType<typeof rootReducer>

export type AppDispatchType = typeof store.dispatch

export const useAppDispatch = () => useDispatch<AppDispatchType>()

//sagaWatcher
sagaMiddleware.run(rootWatcher)

function* rootWatcher() {
  yield takeEvery('BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD', BookingUploadPetImgSagaWorker)
}
