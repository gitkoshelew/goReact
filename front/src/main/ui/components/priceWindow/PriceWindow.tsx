import s from './PriceWindow.module.scss'

export type PriceWindowPropsType = {
    /**
     * Shows price per room
     */
  price: number
}

const { priceWindow } = s

export const PriceWindow = ({ price }: PriceWindowPropsType) => {
  return (
    <div className={priceWindow}>
      <p>${price} night</p>
    </div>
  )
}
