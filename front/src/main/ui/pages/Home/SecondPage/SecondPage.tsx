import React from 'react'
import s from './SecondPage.module.css';
import {StatisticBlockElement} from './statisticBlockElement/StatisticBlockElement';
import {CeoMessText} from './CeoMessText/CeoMessText';
import {CeoMessImg} from './CeoMessImg/CeoMessImg';
import {CeoMessSignature} from './CeoMessSignature/CeoMessSignature';
import {ScrollBlock} from './ScrollBlock/ScrollBlock';

const {secondPage, statisticBlock, ceoMessTitle} = s;

export const SecondPage = () => {
    return (
        <div className={secondPage}>
            <ScrollBlock/>
            <div className={statisticBlock}>
                <StatisticBlockElement statisticNum={20} statisticText={'VARIOUS SERVICES'}/>
                <StatisticBlockElement statisticNum={150} statisticText={'DIFFERENT ROOMS'}/>
                <StatisticBlockElement statisticNum={15} statisticText={'EXPERIENCE'}/>
            </div>
            <div className={ceoMessTitle}>
                <CeoMessImg/>
                <div>
                    <CeoMessText type={'bigBlackText'}/>
                    <CeoMessText type={'smallGreyText'}/>
                </div>
                <CeoMessSignature/>
            </div>
        </div>
    )
}