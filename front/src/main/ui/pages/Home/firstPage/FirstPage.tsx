import React from 'react'
import s from './FirstPage.module.css';
import {HomeTablet} from '../../../components/HomeTablet/HomeTablet';
import {HomeTabletSmall} from '../../../components/HomeTablet/HomeTabletSmall/HomeTabletSmall';

const {homeGreenText, homeBigBlackText, greyText, textBlock, photoBlock, homeTablet,homeTitle,homeSmallTablet} = s;

export const FirstPage = () => {
    return (
        <div className={homeTitle}>
            <div className={textBlock}>
                <div className={homeGreenText}>
                    Welcome home
                </div>
                <div className={homeBigBlackText}>
                    <p>Our world is your</p>
                    <p>playground.</p>
                </div>
                <div className={greyText}>
                    <p>Make yourself at home in our sophisticated guest</p>
                    <p>rooms, take in the incredible views and enjoy fresh air</p>
                    <p>from our beautiful sea city.</p>
                </div>
                <div className={homeTablet}>
                    <HomeTablet/>
                </div>
                <div className={homeSmallTablet}>
                    <HomeTabletSmall/>
                </div>
            </div>
            {/*<div className={photoBlock}>*/}
            {/*    <p>373x335</p>*/}
            {/*</div>*/}
        </div>
    )
}