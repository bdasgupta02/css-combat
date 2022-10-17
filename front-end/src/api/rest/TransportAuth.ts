import { AxiosInstance } from 'axios'

// TODO change pw and forgot pw
class TransportAuth {
  client: AxiosInstance

  constructor(client: AxiosInstance) {
    this.client = client
  }

  signIn(type: string, identifier: string, password: string) {
    return this.client.post('/auth/sign-in', {
      type,
      identifier,
      password,
    })
  }

  signUp(email: string, username: string, fullName: string, password: string) {
    return this.client.post('/auth/sign-up', {
      email,
      username,
      fullName,
      password,
    })
  }
}

export default TransportAuth
