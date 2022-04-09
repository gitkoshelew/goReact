import { apiBookingRegForm } from './API'

type UploadFleRespType = {
  type: string
  ownerId: number
}
type PhotoUrl = {
  photoUrl: string | null
}

export const FilesAPI = {
  async uploadFile(file: File): Promise<PhotoUrl> {
    let formData = new FormData()
    formData.append('image', file, file.name)
    formData.append('type', 'user')
    formData.append('ownerId', '1')
    const res = apiBookingRegForm.post<UploadFleRespType, PhotoUrl>('save', formData)
    console.log(res)
    return res
  },
}
