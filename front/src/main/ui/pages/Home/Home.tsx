import React from 'react'
import s from './Home.module.css';
import {FirstPage} from './firstPage/FirstPage';
import {SecondPage} from './SecondPage/SecondPage';
import {ThirdPage} from './ThirdPage/ThirdPage';

const {secondPage, statisticBlock, statisticBlockNum, statisticBlockText, statisticElem} = s;

export const Home = () => {
    return (
        <div>
            <FirstPage/>
            {/*<SecondPage/>*/}
            {/*<ThirdPage/>*/}
        </div>
    )
}