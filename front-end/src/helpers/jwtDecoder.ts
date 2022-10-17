import jwt_decode from 'jwt-decode'

type JwtPayload = {
  userId: number
  username: string
}

const jwtDecoder = (jwt: string): JwtPayload => {
  const claims: any = jwt_decode(jwt)

  return {
    userId: claims.userId,
    username: claims.username,
  }
}

export default jwtDecoder
