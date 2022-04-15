import { apiBookingRegForm } from './API'
import { AxiosResponse } from 'axios'

export type PhotoUrl = {
  original: string
  qvga: string
  vga: string
  hd720p: string
}

export type FetchPhotoUrlResponse = PhotoUrl

export const FilesAPI = {
  async uploadFile(file: File): Promise<AxiosResponse<FetchPhotoUrlResponse>> {
    let formData = new FormData()
    formData.append('image', file, file.name)
    formData.append('type', 'user')
    formData.append('ownerId', '1')
    const res = apiBookingRegForm.post('save', formData)
    return res
  },
}
