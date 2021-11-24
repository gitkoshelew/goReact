import React from 'react';
import errImg from '../../../../assets/img/404Page/notfound.gif';
import s from './error404.module.css';

const{err404}=s;

export const Error404 = () => {
    return (
        <div className={err404}>
            <img src={errImg} alt="ErrorImg"/>
        </div>
    );
}


