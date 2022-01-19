import * as Yup from 'yup'

const phoneRegExp =
  /^((\\+[1-9]{1,4}[ \\-]*)|(\\([0-9]{2,3}\\)[ \\-]*)|([0-9]{2,4})[ \\-]*)*?[0-9]{3,4}?[ \\-]*[0-9]{3,4}?$/
export const RegistrationSchema = Yup.object().shape({
  password: Yup.string().min(8, 'Too short!').max(20, 'Too long!').required('Required'),
  email: Yup.string().email('Invalid email').required('Required'),
  name: Yup.string().required('Required').min(2, 'Too short!').max(20, 'Too long!'),
  sName: Yup.string().required('Required').min(2, 'Too short!').max(20, 'Too long!'),
  mName: Yup.string().required('Required').min(2, 'Too short!').max(20, 'Too long!'),
  sex: Yup.string().required('Required'),
  birthDate: Yup.string().required('Required'),
  address: Yup.string().required('Required'),
  phone: Yup.string().matches(phoneRegExp, 'Phone number is not valid'),
})
