import s from './PriceWindow.module.css';

type PriceWindowPropsType = {
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