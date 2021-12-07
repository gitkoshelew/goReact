import React from 'react'
import s from './StatisticBlockElement.module.css';

const {statisticBlockNum, statisticBlockText, statisticElem} = s;

type StatisticBlockElementPropsType = {
    statisticNum: number
    statisticText: string

}


export const StatisticBlockElement = (props: StatisticBlockElementPropsType) => {
    const {statisticNum, statisticText} = props;
    return (
        <div className={statisticElem}>
            <div className={statisticBlockNum}>{statisticNum}+</div>
            <div className={statisticBlockText}>{statisticText}</div>
        </div>
    )
}