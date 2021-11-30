import {Home} from '../pages/Home/Home';
import {Navigate, Route, Routes} from 'react-router-dom';
import {Hotels} from '../pages/Hotels/Hotels';
import {AboutUs} from '../pages/AboutUs/AboutUs';
import {Error404} from '../pages/error404/error404';
import {Login} from '../pages/Login/Login';
import {Gallery} from '../pages/Gallery/Gallery';
import {Room} from '../pages/Room/Room';
import {Blog} from '../pages/Blog/Blog';
import {Service} from '../pages/Service/Service';
import {Basket} from '../pages/Basket/Basket';

export const PATH = {
    HOME:'/home',
    LOGIN: '/login',
    HOTELS: '/hotels',
    ABOUT_US: '/aboutus',
    ROOM:'/room',
    SERVICE:'/service',
    BLOG:'/blog',
    GALLERY:'/gallery',
    BASKET:'/basket'


}


export const RoutesInfo = () => {
    return (
       <div>
            <Routes>
                <Route path={'/'} element={<Navigate replace to={PATH.HOME} />} />
                <Route path={PATH.HOME} element={<Home/>}/>
                <Route path={PATH.HOTELS} element={<Hotels/>}/>
                <Route path={PATH.ABOUT_US} element={<AboutUs/>}/>
                <Route path={PATH.LOGIN} element={<Login/>}/>
                <Route path={PATH.GALLERY} element={<Gallery/>}/>
                <Route path={PATH.ROOM} element={<Room/>}/>
                <Route path={PATH.BLOG} element={<Blog/>}/>
                <Route path={PATH.SERVICE} element={<Service/>}/>
                <Route path={PATH.BASKET} element={<Basket/>}/>

                <Route path={"*"} element={<Error404/>}/>
            </Routes>
       </div>


    )
}