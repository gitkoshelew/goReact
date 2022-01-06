import axios from 'axios'

const API_URL = 'http://localhost:8080/'

const settings = {
  withCredentials: true,
  baseURL: API_URL,
}

export const $api = axios.create({
  ...settings,
})

// $api.interceptors.request.use((config) => {
//     console.log(config)
// })

export const logIn = async (data: any) => {
  const promise = await $api.post('login', { data })
  return promise
}
