import t from '../../../../assets/img/navBar/t.svg'
import f from '../../../../assets/img/navBar/Vector.svg'
import dot from '../../../../assets/img/navBar/dot.svg'
import s from './Logo.module.css';
import React from 'react';

const{logo}=s;

export const Logo = () => {
    return (
        <div className={logo}>
            <img src={f} alt="f"/>
            <img src={t} alt="t"/>
            <img src={dot} alt="dot"/>
        </div>
    )
}