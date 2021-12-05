import { t } from '../../svgWrapper/navBarSvgWrapper'
import { f } from '../../svgWrapper/navBarSvgWrapper'
import { dot } from '../../svgWrapper/navBarSvgWrapper'
import s from './Logo.module.css';
import React from 'react';

const { logo } = s;

export const Logo = () => {
    return (
        <div className={logo}>
            <img src={f} alt="f"/>
            <img src={t} alt="t"/>
            <img src={dot} alt="dot"/>
        </div>
    )
}