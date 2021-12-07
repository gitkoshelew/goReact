import {HotelsSaga} from './HotelsSaga';


export function* rootWatcher() {
    yield HotelsSaga()
}