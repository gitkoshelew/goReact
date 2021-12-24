import React from 'react'
import s from './CeoMessText.module.scss'

const { mainText, secondaryText } = s

type CeoMessTextPropsType = {
  type: string
}

export const CeoMessText = (props: CeoMessTextPropsType) => {
  const { type } = props
  return (
    <>
      {type === 'mainText' && (
        <div className={mainText}>
          <p>Beach Hotel More than a stay</p>
        </div>
      )}
      {type === 'secondaryText' && (
        <div className={secondaryText}>
          <p>We have a lot of effort to bring more quality</p>
          <p>time to you and the people you love. You will</p>
          <p>have a chance to enjoy meaningful moments</p>
          <p>together and that's reason why we're here.</p>
        </div>
      )}
    </>
  )
}
