import createSagaMiddleware from 'redux-saga'
import {combineReducers} from 'redux'
import {configureStore} from '@reduxjs/toolkit';
import {hotelsReducer} from '../redusers/Hotels-reducer';
import {rootWatcher} from '../saga/rootSaga';


const sagaMiddleWare = createSagaMiddleware();


export const rootReducer = combineReducers({
    hotels: hotelsReducer
})

export type RootReducerType = typeof rootReducer
export type AppRootStateType = ReturnType<RootReducerType>


export const store = configureStore({
    reducer: rootReducer,
    middleware: [sagaMiddleWare]
});


sagaMiddleWare.run(rootWatcher)

