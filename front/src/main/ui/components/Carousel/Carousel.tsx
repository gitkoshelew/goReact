import React, { ReactChild, ReactChildren } from 'react'
import Slider, { ResponsiveObject } from 'react-slick'
import s from './Carousel.module.scss'
import { nextArrow, prevArrow, sliderDot } from '../../svgWrapper/HomeSvgWrapper'

const { carousel, carouselBig } = s

export type CarouselPropsType = {
  children: ReactChild | ReactChild[] | ReactChildren | ReactChildren[]
  /**
   *  2 breakpoints (1250, 768) with settings:
   *
   *  slidesToShow
   *
   *  slidesToScroll
   *
   *  infinite
   *
   *  dots
   *
   *  arrows
   */
  responsiveArrForCarousel?: ResponsiveObject[]
  startSlideToShow: number
  startSlideScroll: number
  /**
   * The display content of the Carousel:
   *
   * smallWidth
   *
   * bigWidth
   */
  type: string
  /**
   * Show arrows or not
   */
  isWithArrows: boolean
}
export type SampleArrowPropsType = {
  className?: string
  onClick?: () => void
}

export const Carousel = ({
  children,
  responsiveArrForCarousel,
  startSlideToShow,
  startSlideScroll,
  type,
  isWithArrows,
}: CarouselPropsType) => {
  function SampleNextArrow(props: SampleArrowPropsType) {
    const { className, onClick } = props
    return (
      <div className={className} onClick={onClick}>
        <img src={nextArrow} alt="nextArrow" />
      </div>
    )
  }

  function SamplePrevArrow(props: SampleArrowPropsType) {
    const { className, onClick } = props
    return (
      <div className={className} onClick={onClick}>
        <img src={prevArrow} alt="nextArrow" />
      </div>
    )
  }

  const settings = {
    customPaging: function () {
      return (
        <button>
          <img src={sliderDot} alt={'sliderDot'} />
        </button>
      )
    },
    dotsClass: `slick-dots slick-thumb `,
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: startSlideToShow,
    slidesToScroll: startSlideScroll,
    arrows: isWithArrows,
    responsive: responsiveArrForCarousel,
    nextArrow: <SampleNextArrow />,
    prevArrow: <SamplePrevArrow />,
  }

  return (
    <>
      {type === 'smallWidth' && (
        <div className={carousel}>
          <Slider {...settings}>{children}</Slider>
        </div>
      )}
      {type === 'bigWidth' && (
        <div className={carouselBig}>
          <Slider {...settings}>{children}</Slider>
        </div>
      )}
    </>
  )
}
