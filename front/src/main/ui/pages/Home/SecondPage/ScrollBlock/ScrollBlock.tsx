import React from 'react'
import s from './ScrollBlock.module.scss'
import arrow from '../../../../../../assets/img/home/secondPage/Arrow.svg'

const { scrollBlock } = s

export const ScrollBlock = () => {
  return (
    <div className={scrollBlock}>
      <p>Scroll</p>
      <img src={arrow} alt="arrowScroll" />
    </div>
  )
}
