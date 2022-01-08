import s from './Login.module.scss'
import {Form, Formik} from 'formik'
import {authenticationSchema} from './validations/authValidation'
import {TextField} from '../../components/TextField/TextField'
import {AuthAPI, UserType} from '../../../dal/api_client/API';

const {authenticationForm, authenticationTitle, sendReqBtn, sendReqErrorBtn} = s

export const LoginPage = () => {
    return (
        <Formik
            initialValues={{
                email: '',
                password: '',
            }}
            validationSchema={authenticationSchema}
            onSubmit={(user: UserType) => AuthAPI.logIn(user).then(console.log)}
        >
            {(formik) => (
                <Form className={authenticationForm}>
                    <div className={authenticationTitle}>Log in</div>
                    <TextField label="Email" name="email" type="text"/>
                    <TextField label="Password" name="password" type="text"/>
                    <button
                        className={formik.errors.email || formik.errors.password ? sendReqErrorBtn : sendReqBtn}
                        type="submit"
                    >
                        Login
                    </button>
                </Form>
            )}
        </Formik>
    )
}
