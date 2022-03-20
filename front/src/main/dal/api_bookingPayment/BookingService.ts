import { AxiosResponse } from 'axios'
import { apiBookingPayment } from './API'

export type BookingPaymentFormType = {
  firstName: string
  lastName: string
  email: string
  cardNumber: string
  company: string
  mm: string
  yy: string
  cvv: string
}
export type FetchBookingPaymentResponse = BookingPaymentFormType[]

export const BookingPageAPI = {
  getAllBookingPayment(): Promise<AxiosResponse<FetchBookingPaymentResponse>> {
    return apiBookingPayment.get('api/cards')
  },
  async createBookingPayment(newCard: BookingPaymentFormType): Promise<AxiosResponse<FetchBookingPaymentResponse>> {
    const res = await apiBookingPayment.post('api/cards', newCard)
    console.log(res)
    return res
  },
}
