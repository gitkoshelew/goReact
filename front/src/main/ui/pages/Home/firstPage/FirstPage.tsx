import React from 'react'
import s from './FirstPage.module.css';
import {HomeTablet} from '../../../components/HomeTablet/HomeTablet';
import {HomeTabletSmall} from '../../../components/HomeTablet/HomeTabletSmall/HomeTabletSmall';

const {
    firstPageLinkText,
    firstPageMainText,
    secondaryText,
    textBlock,
    photoBlock,
    homeTablet,
    firstPageTitle,
    homeSmallTablet,
    photoTabletBlock,
    informText
} = s;

export const FirstPage = () => {
    return (
        <div className={firstPageTitle}>
            <div className={textBlock}>
                <div className={informText}>
                    <div className={firstPageLinkText}>
                        Welcome home
                    </div>
                    <div className={firstPageMainText}>
                        <p>Our world is your</p>
                        <p>playground.</p>
                    </div>
                    <div className={secondaryText}>
                        <p>Make yourself at home in our sophisticated guest</p>
                        <p>rooms, take in the incredible views and enjoy fresh air</p>
                        <p>from our beautiful sea city.</p>
                    </div>
                </div>
                <div className={photoTabletBlock}>
                    <div className={photoBlock}>
                        <p>373x335</p>
                    </div>
                    <div className={homeTablet}>
                        <HomeTablet/>
                    </div>
                </div>
                <div className={homeSmallTablet}>
                    <HomeTabletSmall/>
                </div>
            </div>

        </div>
    )
}