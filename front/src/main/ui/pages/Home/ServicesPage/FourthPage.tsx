import s from './FourthPage.module.scss'
import { TitlePageTextBlock } from '../../../components/TitlePageTextBlock/TitlePageTextBlock'
import { Carousel } from '../../../components/Carousel/Carousel'
import { ServiceBlock } from './ServiceBlock/ServiceBlock'
import React from 'react'

const { fourthPage, titleText, oneSlide, serviceCarouselBlock } = s
export const FourthPage = () => {
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

  return (
    <div className={fourthPage}>
      <div className={titleText}>
        <TitlePageTextBlock isWithLink={true} linkTextMess={'See more'} mainTextMess={'services'} />
      </div>
      <div className={serviceCarouselBlock}>
        <ServiceBlock />
        <Carousel
          isWithArrows={true}
          type={'smallWidth'}
          startSlideScroll={startSlideScroll}
          startSlideToShow={startSlideToShow}
          responsiveArrForCarousel={responsiveArrForCarousel}
        >
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
    </div>
  )
}
