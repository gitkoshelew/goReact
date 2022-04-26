import { apiAuthSocialNetwork } from './API'

export const SocialNetworkAuthAPI = {
  getLinkedinLogin(): Promise<void> {
    return apiAuthSocialNetwork.get('api/linkedinlogin')
  },
  getGithubAuth(): Promise<void> {
    return apiAuthSocialNetwork.get('api/gitlogin')
  },
}
