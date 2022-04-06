import { AxiosResponse } from 'axios'
import { apiBookingRegForm } from './API'

export type BookingInfoPhoto = {
  type: string
  ownerId: number
}

export type BookingPhotoType = {
  photoUrl: null | string
}

export type FetchBookingPhotoResponse = BookingInfoPhoto[] & BookingPhotoType[]

export const BookingRegFormAPI = {
  async createBookingPhoto(
    newInfoPhoto: BookingInfoPhoto,
    newPhoto: BookingPhotoType
  ): Promise<AxiosResponse<FetchBookingPhotoResponse>> {
    const res = await apiBookingRegForm.post('save', { newInfoPhoto: newInfoPhoto, newPhoto: newPhoto })
    return res
  },
}
