import React from 'react'
import s from './CeoMessSignature.module.css';
import signature from '../../../../../../assets/img/home/secondPage/signature.svg'

const {ceoMessSignature} = s;

export const CeoMessSignature = () => {
    return (

        <div className={ceoMessSignature}>
            <img src={signature} alt="signature"/>
            <p>CEO Alex</p>
        </div>
    )
}