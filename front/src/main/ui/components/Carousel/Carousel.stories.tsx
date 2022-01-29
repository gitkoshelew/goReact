import React from 'react';
import { Story, Meta } from '@storybook/react';
import s from './Carousel.module.scss'
import { Carousel, CarouselPropsType, SampleArrowPropsType } from './Carousel';

const { oneSlide } = s

const meta: Meta = {
    title: 'Carousel',
    component: Carousel,
}
export default meta
const startSlideToShow = 1
const startSlideScroll = 1

const responsiveArrForCarousel = [
    {
        breakpoint: 1250,
        settings: {
            slidesToShow: 2,
            slidesToScroll: 1,
            infinite: true,
            dots: true,
            arrows: false,
        },
    },
    {
        breakpoint: 768,
        settings: {
            slidesToShow: 1,
            slidesToScroll: 1,
            infinite: true,
            dots: true,
            arrows: false,
        },
    },
]

export const Template: Story<CarouselPropsType & SampleArrowPropsType> = (args) => <Carousel {...args}>
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


export const SmallWidth = Template.bind({});
export const BigWidth = Template.bind({});

SmallWidth.args = {
    type: 'smallWidth',
    isWithArrows: true,
    startSlideScroll: startSlideToShow,
    startSlideToShow: startSlideScroll,
    responsiveArrForCarousel: responsiveArrForCarousel

}
BigWidth.args = {
    type: 'bigWidth',
    isWithArrows: true,
    startSlideScroll: startSlideToShow,
    startSlideToShow: startSlideScroll,
    responsiveArrForCarousel: responsiveArrForCarousel
}

