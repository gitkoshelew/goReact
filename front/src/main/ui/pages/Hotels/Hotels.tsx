import {useDispatch, useSelector} from 'react-redux';
import {AppRootStateType} from '../../../bll/store/store';
import {useEffect} from 'react';
import {fetchHotelsAC} from '../../../bll/saga/HotelsSaga';
import s from './Hotels.module.css';

type fakeUserType = {
    avatar: string
    email: string
    first_name: string
    id: number
    last_name: string
}

const {hotelsTable} = s;

export const Hotels = () => {


    const hotels = useSelector<AppRootStateType, fakeUserType[]>(state => state.hotels.hotels)
    const dispatch = useDispatch()


    useEffect(() => {
        dispatch(fetchHotelsAC())
    }, [])


    return (<div>
        {hotels.map(t => <div className={hotelsTable}>
            id:{t.id}
            name:{`${t.first_name} ${t.last_name}`}
            email:{t.email}
            photo:
            <img src={t.avatar} alt="photo"/>

        </div>)}
    </div>)
}