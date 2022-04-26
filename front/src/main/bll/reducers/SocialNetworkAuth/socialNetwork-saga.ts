import { call, put } from 'redux-saga/effects'
import { socialNetworkAuthStart, socialNetworkAuthSuccess } from './socialNetwork-reducer'
import { SocialNetworkAuthAPI } from '../../../dal/api_socialNetwork/SocialNetworkAuthService'

export function* fetchGithubAuthSagaWorker(action: ReturnType<typeof fetchGithubRequest>) {
  yield put(socialNetworkAuthStart())
  yield call(SocialNetworkAuthAPI.getGithubAuth)
  yield put(socialNetworkAuthSuccess())
}

export const fetchGithubRequest = () => ({
  type: 'SOCIAL_NETWORK/GITHUB_SAGA',
})

export function* fetchLinkedinAuthSagaWorker(action: ReturnType<typeof fetchLinkedinRequest>) {
  yield put(socialNetworkAuthStart())
  yield call(SocialNetworkAuthAPI.getLinkedinLogin)
  yield put(socialNetworkAuthSuccess())
}

export const fetchLinkedinRequest = () => ({
  type: 'SOCIAL_NETWORK/Linkedin_SAGA',
})
