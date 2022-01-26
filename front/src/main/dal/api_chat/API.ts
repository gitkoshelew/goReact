import axios from 'axios'

export const API_CHAT_URL = process.env.REACT_APP_API_CHAT_LINK

export const apiChat = axios.create({
  baseURL: API_CHAT_URL,
})
