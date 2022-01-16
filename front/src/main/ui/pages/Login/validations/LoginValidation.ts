import * as Yup from 'yup'

export const authenticationSchema = Yup.object().shape({
  password: Yup.string().min(8, 'Too short!').max(20, 'Too long!').required('Required'),
  email: Yup.string().email('Invalid email').required('Required'),
})
