import React from 'react'

import s from './Select.module.css'

const { selectContainer } = s

export const SelectUI = () => {
  return (
    <div className={selectContainer}>
      <select>
        <option value={'1'}>EN</option>
        <option value={'2'}>RU</option>
        <option value={'3'}>FR</option>
      </select>
    </div>
  )
}
