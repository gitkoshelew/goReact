/* eslint-disable */
import React from 'react'
import { ErrorMessage, useField } from 'formik'
import s from './TextField.module.scss'

const { inputTitle, inputField, inputFieldError, inputLabel, errorMsg } = s

export const TextField = ({ label, ...props }: any) => {
  const [field, meta] = useField(props)
  return (
    <div className={inputTitle}>
      <label htmlFor={field.name} className={inputLabel}>
        {label}
      </label>
      <input className={meta.touched && meta.error ? inputFieldError : inputField} {...field} />
      <ErrorMessage component="div" name={field.name} className={errorMsg} />
    </div>
  )
}
