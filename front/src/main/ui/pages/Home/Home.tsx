import React from 'react'
import { FirstPage } from './firstPage/FirstPage'
import { SecondPage } from './SecondPage/SecondPage'
import { ThirdPage } from './ThirdPage/ThirdPage'
import { FourthPage } from './FourthPage/FourthPage'
import { FifthPage } from './FifthPage/FifthPage'
import { SixthPage } from './SixthPage/SixthPage'

export const Home = () => {
  console.log(process.env)
  return (
    <div>
      <FirstPage />
      <SecondPage />
      <ThirdPage />
      <FourthPage />
      <FifthPage />
      <SixthPage />
    </div>
  )
}
