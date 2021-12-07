import React from 'react';
// @ts-ignore
import Slider from 'react-slick';
import s from './Carousel.module.css';
import {sliderDot} from '../../svgWrapper/HomeSvgWrapper';


const {carousel, oneSlide} = s;

export const Carousel = ({children}:any) => {
    const settings = {
        customPaging: function (i: any) {
            return (
                <a>
                    <img src={sliderDot} alt={'sliderDot'}/>
                </a>

            );
        },
        dotsClass: `slick-dots slick-thumb `,
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1,
        arrows: false,
    }

    return (
        <div className={carousel}>
            <Slider {...settings}>
                {children}
            </Slider>
        </div>
    );
}