import { TitlePageTextBlock } from '../../../components/TitlePageTextBlock/TitlePageTextBlock'
import s from './SixthPage.module.scss'
import { FeedBack } from '../../../components/FeedBack/FeedBack'
import { flower, forkAndSpoon, swimmer } from '../../../svgWrapper/HomeSvgWrapper'
import { Carousel } from '../../../components/Carousel/Carousel'

const { sixthPage, sixthPageTitle } = s

export const SixthPage = () => {
  const startSlideToShow = 3
  const startSlideScroll = 3

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
    <div className={sixthPage}>
      <div className={sixthPageTitle}>
        <TitlePageTextBlock mainTextMess={'our guests love us'} isWithLink={false} />
      </div>
      <Carousel
        isWithArrows={false}
        type={'bigWidth'}
        startSlideScroll={startSlideScroll}
        startSlideToShow={startSlideToShow}
        responsiveArrForCarousel={responsiveArrForCarousel}
      >
        <FeedBack
          feedBackMess={`The staff have been amazing and extremely helpful.
                         They respond in a very friendly manner to all
                          questions of us and we will comeback if we have a chance`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={flower}
        />
        <FeedBack
          feedBackMess={`Every year, we come back to FT hotel for our holiday. 
                     It was extremely interesting and enjoyable !
                      We are happy with our stay in this hotel and we love meals here.`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={swimmer}
        />
        <FeedBack
          feedBackMess={`The location of FT hotel is perfect and very central.
                 I was very happy to walk around the hotel, discover surrounding scenes such as beach, night market, museum`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={forkAndSpoon}
        />
        <FeedBack
          feedBackMess={`The staff have been amazing and extremely helpful.
                They respond in a very friendly manner to all questions of us and we will comeback if we have a chance`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={flower}
        />
        <FeedBack
          feedBackMess={`Every year, we come back to FT hotel for our holiday. 
                     It was extremely interesting and enjoyable ! We are happy with our stay in this hotel and we love meals here.`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={swimmer}
        />
        <FeedBack
          feedBackMess={`The location of FT hotel is perfect and very central.
                 I was very happy to walk around the hotel, discover surrounding scenes such as beach, night market, museum`}
          nationality={'American'}
          userName={'Ralph Edwards'}
          photo={forkAndSpoon}
        />
      </Carousel>
    </div>
  )
}
