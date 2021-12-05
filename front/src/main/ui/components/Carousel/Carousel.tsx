import React from 'react';
// @ts-ignore
import Slider from 'react-slick';
import s from './Carousel.module.css';
import {nextArrow, prevArrow, sliderDot} from '../../svgWrapper/HomeSvgWrapper';


const {carousel,carouselBig} = s;

type CarouselPropsType = {
    children: any
    responsiveArrForCarousel?: any
    startSlideToShow: number
    startSlideScroll: number
    type: string
    isWithArrows:boolean
}


export const Carousel = ({
                             children,
                             responsiveArrForCarousel,
                             startSlideToShow,
                             startSlideScroll,
                             type,
                             isWithArrows
                         }: CarouselPropsType) => {

    function SampleNextArrow(props: any) {
        const {className, onClick} = props;
        return (
            <div
                className={className}
                // style={{ ...style, display: "block", background: "red" }} to custom update
                onClick={onClick}
            >
                <img src={nextArrow} alt="nextArrow"/>
            </div>
        );
    }

    function SamplePrevArrow(props: any) {
        const {className, onClick} = props;
        return (
            <div
                className={className}
                // style={{ ...style, display: "block", background: "red" }} to custom update
                onClick={onClick}
            >
                <img src={prevArrow} alt="nextArrow"/>
            </div>
        );
    }


    const settings = {
        customPaging: function (i: any) {
            return (
                <button>
                    <img src={sliderDot} alt={'sliderDot'}/>
                </button>

            );
        },
        dotsClass: `slick-dots slick-thumb `,
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: startSlideToShow,
        slidesToScroll: startSlideScroll,
        arrows: isWithArrows,
        responsive: responsiveArrForCarousel,
        nextArrow: <SampleNextArrow/>,
        prevArrow: <SamplePrevArrow/>
    }

    return (
        <>
            {type === 'smallWidth' && <div className={carousel}>
                <Slider {...settings}>
                    {children}
                </Slider>
            </div>
            }
            {type === 'bigWidth' &&
            <div className={carouselBig}>
                <Slider {...settings}>
                    {children}
                </Slider>
            </div>
            }
        </>
    );
}