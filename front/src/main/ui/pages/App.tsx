import React, {useState} from 'react';
import './App.css';
import {NavBar} from './navBar/navBar';
import {RoutesInfo} from '../Routes/RoutesInfo';
import {Footer} from './footer/footer';


function App() {

    const [isBurgerCollapse, setIsBurgerCollapse] = useState(false);


    return (
        <div>
            <NavBar setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse}/>
            <div>
                <RoutesInfo/>
                {/*<Footer/>*/}
            </div>
        </div>
    );
}

export default App;
