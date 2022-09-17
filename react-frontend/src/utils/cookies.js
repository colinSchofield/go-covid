import Cookies from 'universal-cookie'

const cookies = new Cookies()
const COOKIE_KEY = "covid-19-userId"
const SESSION_KEY = "covid-19-session"

export const setUserIdCookie = (value) => {
  cookies.set(COOKIE_KEY, value, { path: '/', maxAge: Number.MAX_SAFE_INTEGER })
}

export const getUserIdCookie = () => {
  return cookies.get(COOKIE_KEY)
}

export const clearUserIdCookie = () => {
  cookies.set(COOKIE_KEY, null, { path: '/' })
}

export const setSessionCookie = () => {
  cookies.set(SESSION_KEY, "Don't display the geo-located graph again this session", { path: '/' })
}

export const isSessionCookieSet = () => {
  return (cookies.get(SESSION_KEY))
}