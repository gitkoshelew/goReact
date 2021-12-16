import { combineReducers } from "redux";
import { configureStore } from "@reduxjs/toolkit";
import createSagaMiddleware from "redux-saga";
import { BookingRegFormReducer } from "../reducers/BookingRegForm-reducer";
import { useDispatch } from "react-redux";

const sagaMiddleware = createSagaMiddleware();

const rootReducer = combineReducers({
  BookingRegForm: BookingRegFormReducer
});

export type RootReducerType = typeof rootReducer


export const store = configureStore({
  reducer: rootReducer,
  middleware: getDefaultMiddleware => getDefaultMiddleware().prepend(sagaMiddleware)
});


export type AppRootStateType = ReturnType<typeof rootReducer>

export type AppDispatchType=typeof store.dispatch

export const useAppDispatch = () => useDispatch<AppDispatchType>()

