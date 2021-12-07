import React from 'react';
import './App.css';
import {NavBar} from '../components/navBar/navBar';
import {RoutesInfo} from '../Routes/RoutesInfo';
import {Footer} from './footer/footer';
// @ts-ignore
import Rotate from 'react-reveal/Fade'


function App() {
    return (
        <div>
            <Rotate bottom left>
            <NavBar/>
            </Rotate>
            <Rotate top right>
            <RoutesInfo/>
            </Rotate>
            <Footer/>
        </div>
    );
}

export default App;
