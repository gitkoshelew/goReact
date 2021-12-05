import React from 'react';


import s from './Select.module.css';

const { selectContainer } = s

export const SelectUI = () => {


    return (
        <div className={selectContainer}>
            <select>
                <option>EN</option>
                <option>RU</option>
                <option>FR</option>
            </select>
        </div>
    );
}
