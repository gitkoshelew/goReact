import {StartPage} from '../pages/StartPage/StartPage';
import {Navigate, Route, Routes} from 'react-router-dom';
import {Hotels} from '../pages/Hotels/Hotels';
import {AboutUs} from '../pages/AboutUs/AboutUs';
import {Error404} from '../pages/error404/error404';
import {Login} from '../pages/Login/Login';

export const PATH = {
    HOME:'/home',
    LOGIN: '/login',
    HOTELS: '/hotels',
    ABOUT_US: '/aboutUs'
}


export const RoutesInfo = () => {
    return (
       <div>
            <Routes>
                <Route path={'/'} element={<Navigate replace to={PATH.HOME} />} />
                <Route path={PATH.HOME} element={<StartPage/>}/>
                <Route path={PATH.HOTELS} element={<Hotels/>}/>
                <Route path={PATH.ABOUT_US} element={<AboutUs/>}/>
                <Route path={PATH.LOGIN} element={<Login/>}/>

                <Route path={"*"} element={<Error404/>}/>
            </Routes>
       </div>


    )
}