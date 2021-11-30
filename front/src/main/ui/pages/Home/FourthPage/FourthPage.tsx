import s from './FourthPage.module.css';
import {TitlePageTextBlock} from '../../../components/TitlePageTextBlock/TitlePageTextBlock';
import {Carousel} from '../../../components/Carousel/Carousel';
import {ServiceBlock} from './ServiceBlock/ServiceBlock';
import React from 'react';


const {fourthPage, titleText, oneSlide} = s;
export const FourthPage = () => {
    return (
        <div className={fourthPage}>
            <div className={titleText}>
                <TitlePageTextBlock isWithLink={true} linkTextMess={'See more'} mainTextMess={'services'}/>
            </div>
            <ServiceBlock/>
            <Carousel>
                <div className={oneSlide}>
                    <p>photo</p>
                </div>
                <div className={oneSlide}>
                    <p>photo</p>
                </div>
                <div className={oneSlide}>
                    <p>photo</p>
                </div>
                <div className={oneSlide}>
                    <p>photo</p>
                </div>
            </Carousel>
        </div>
    )
}