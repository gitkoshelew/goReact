/* eslint-disable */
import React from 'react'
import { ErrorMessage, useField } from 'formik'
import s from './TextField.module.scss'

const {
  inputTitle,
  inputField,
  inputFieldError,
  inputLabel,
  error__msg,
  registerDataField,
  inputRegisterField,
  inputRegisterFieldError,
  inputRegisterTitle,
  registerField_error__msg,
} = s

export const TextField = ({ label, inputType, inputMsgLabel, ...props }: any) => {
  const [field, meta] = useField(props)
  return (
    <>
      {inputType === 'login' && (
        <div className={inputTitle}>
          <label htmlFor={field.name} className={inputLabel}>
            {label}
          </label>
          <input
            autocomplete="new-password"
            className={meta.touched && meta.error ? inputFieldError : inputField}
            {...field}
            {...props}
          />
          <ErrorMessage component="div" name={field.name} className={error__msg} />
        </div>
      )}
      {inputType === 'register' && (
        <div className={inputRegisterTitle}>
          <div className={registerDataField}>
            <div>{inputMsgLabel}</div>
            <input
              autocomplete="new-password"
              className={meta.touched && meta.error ? inputRegisterFieldError : inputRegisterField}
              {...field}
              {...props}
            />
          </div>
          <ErrorMessage component="div" name={field.name} className={registerField_error__msg} />
        </div>
      )}
    </>
  )
}
