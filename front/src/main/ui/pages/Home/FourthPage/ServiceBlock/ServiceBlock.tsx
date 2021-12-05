import s from './ServiceBlock.module.css';
import { ServicesElement } from '../../../../components/ServicesElement/ServicesElement';
import { useState } from 'react';

const { serviceBlock,toRows } = s;

export type IsActiveServiceElementType = {
    elem1: boolean
    elem2: boolean
    elem3: boolean
    elem4: boolean
}

export const ServiceBlock = () => {


    const [isActive, setIsActive] = useState<IsActiveServiceElementType>({
        elem1: false,
        elem2: false,
        elem3: false,
        elem4: false,
    })


    return (
        <div className={serviceBlock}>
            <div className={toRows}>
            <ServicesElement setIsActive={setIsActive} isActive={isActive} mainTextMess={'TOP RESTAURANT'}
                             secondaryTextMess={'Breakfast & Dinner'}
                             type={'forkAndSpoon'}/>
            <ServicesElement setIsActive={setIsActive} isActive={isActive} mainTextMess={'BEST SUITES'}
                             secondaryTextMess={'Cool View'}
                             type={'case'}/>
            </div>
            <div className={toRows}>
            <ServicesElement setIsActive={setIsActive} isActive={isActive} mainTextMess={'SPA & WELLNESS'}
                             secondaryTextMess={'Open Daily'}
                             type={'flower'}/>
            <ServicesElement setIsActive={setIsActive} isActive={isActive} mainTextMess={'Swimming Pool'}
                             secondaryTextMess={'Open Daily'}
                             type={'swimmer'}/>
        </div>
        </div>
    )
}