import s from './Button.module.scss'
import { ButtonHTMLAttributes } from 'react'

const { uploadPetPhoto, uploadPetPhotoActive, orderBtn } = s

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  view: string
  isActive?: boolean
  onClick?: () => void
}

export const Button = ({ onClick, view, isActive, ...props }: ButtonProps) => {
  return (
    <div>
      {view === 'order' && (
        <button type={'button'} onClick={onClick} className={orderBtn} {...props}>
          order
        </button>
      )}
      {view === 'upload' && (
        <button type={'submit'} className={isActive ? uploadPetPhotoActive : uploadPetPhoto} {...props}>
          Confirm
        </button>
      )}
    </div>
  )
}

