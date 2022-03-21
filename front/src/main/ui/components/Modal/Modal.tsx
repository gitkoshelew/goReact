import React from 'react'
import s from './Modal.module.scss'

const { modal, content, modalActive, contentActive } = s
type ModalPropsType = {
  active?: boolean
  setActive?: (value: boolean) => void
}
export const Modal: React.FC<ModalPropsType> = ({ active, setActive, children }) => {
  const setActiveHandler = () => {
    if (setActive) {
      setActive(false)
    }
  }
  const contentHandler = (e: React.MouseEvent<HTMLDivElement>) => {
    e.stopPropagation()
  }
  return (
    <div className={active ? modalActive : modal} onClick={setActiveHandler}>
      <div className={active ? contentActive : content} onClick={contentHandler}>
        {children}
      </div>
    </div>
  )
}
