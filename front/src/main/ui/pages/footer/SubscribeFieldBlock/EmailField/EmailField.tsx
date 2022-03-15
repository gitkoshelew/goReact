import s from './EmailField.module.scss'

const { formWrapper } = s

export const EmailField = () => {
  return (
    <form className={formWrapper}>
      <input type="text" placeholder="Enter your email" required />
      <button type="submit">Subscribe</button>
    </form>
  )
}
