import React from 'react'
import s from './CeoMessText.module.css';

const {bigBlackText, smallGreyText} = s;

type CeoMessTextPropsType = {
    type: string
}


export const CeoMessText = (props: CeoMessTextPropsType) => {
    const {type} = props;
    return (
        <>
            {type === 'bigBlackText' &&
            <div className={bigBlackText}>
                <p>Beach Hotel More than a stay</p>
            </div>}
            {type === 'smallGreyText' &&
            <div className={smallGreyText}>
                <p>We have a lot of effort to bring more quality</p>
                <p>time to you and the people you love. You will</p>
                <p>have a chance to enjoy meaningful moments</p>
                <p>together and that's reason why we're here.</p>
            </div>
            }

        </>
    )
}