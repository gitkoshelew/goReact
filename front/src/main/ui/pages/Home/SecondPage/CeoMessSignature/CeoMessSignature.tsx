import React from 'react'
import s from './CeoMessSignature.module.css'
import { signature } from '../../../../svgWrapper/HomeSvgWrapper'

const { ceoMessSignature } = s

export const CeoMessSignature = () => {
  return (
    <div className={ceoMessSignature}>
      <img src={signature} alt="signature" />
      <p>CEO Alex</p>
    </div>
  )
}
