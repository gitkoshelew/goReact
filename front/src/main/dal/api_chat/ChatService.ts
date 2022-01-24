import { AxiosResponse } from 'axios'
import { $api } from '../api_client/API'
import { $apiChat } from './API'

export type ChatUser = {
  userId: number
  email: string
  role: string
  verified: boolean
  name: string
  sName: string
  mName: string
  sex: string
  birthDate: string
  address: string
  phone: string
  photo: string
}
export type ChatMessage = {
  id: string
  senderId: number
  receiverId: number
  conversationId: number
  text: string
}
export type ChatConversation = {
  id: number
  members: number[]
}

export type FetchUsersResponse = ChatUser[]
export type FetchInitMessagesResponse = ChatMessage[]
export type GetConversationResponse = ChatConversation
export type CreateConversationResponse = ChatConversation

export const ChatAPI = {
  async fetchUsers(): Promise<AxiosResponse<FetchUsersResponse>> {
    const res = await $api.get('api/users')
    return res
  },
  async fetchInitMessages(conversationId: number): Promise<AxiosResponse<FetchInitMessagesResponse>> {
    const res = await $apiChat.get(`api/messages/${conversationId}`)
    return res
  },
  async getConversation(firstUserId: number, secondUserId: number): Promise<AxiosResponse<GetConversationResponse>> {
    const res = await $apiChat.get(`api/conversations/find/${firstUserId}/${secondUserId}`)
    return res
  },
  async createConversation(senderId: number, receiverId: number): Promise<AxiosResponse<CreateConversationResponse>> {
    const res = await $apiChat.post(`api/conversations/add`, { senderId, receiverId })
    return res
  },
}
