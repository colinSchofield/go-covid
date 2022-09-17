import { getEndPoint } from './config'

// The End Point is dependent on the generated URL of the Lambda API Gateway (i.e. see deploy.sh!)
const END_POINT = getEndPoint()

const setupHeaders = () => {

  const username = "user"
  const password = "user"
  let headers = new Headers()
  headers.append("Authorization", "Basic " + btoa(username + ":" + password))
  headers.append("Content-Type", "application/json")
  return headers
}

export const getCovid19Daily = () => {

  const url = END_POINT + "/api/1.0/list/daily"

  return fetch(url, { method: "GET", headers: setupHeaders() })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        throw new Error("Network Error, please wait a while and try again: " + response.statusText)
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      throw error
    })
}

export const getCovid19Monthly = (region) => {

  const url = END_POINT + "/api/1.0/list/monthly/" + region

  return fetch(url, { method: "GET", headers: setupHeaders() })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        throw new Error("Network Error, please wait a while and try again: " + response.statusText)
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      throw error
    })
}

export const getRegions = () => {

  const url = END_POINT + "/api/1.0/list/regions"

  return fetch(url, { method: "GET", headers: setupHeaders() })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        throw new Error("Network Error, please wait a while and try again: " + response.statusText)
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      throw error
    })
}

export const createUser = (jsonBody) => {

  const url = END_POINT + "/api/1.0/user"

  return fetch(url, {
          method: "POST",
          headers: setupHeaders(),
          body: JSON.stringify(jsonBody)
        })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        throw new Error("Network Error, please wait a while and try again: " + response.statusText)
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      throw error
    })
}

export const updateUser = (jsonBody) => {

  const url = END_POINT + "/api/1.0/user/" + jsonBody.id

  return fetch(url, {
          method: "PUT",
          headers: setupHeaders(),
          body: JSON.stringify(jsonBody)
        })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        throw new Error("Network Error, please wait a while and try again: " + response.statusText)
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      throw error
    })
}

export const getUser = (userId) => {
  const url = END_POINT + "/api/1.0/user/" + userId

  return fetch(url, {
          method: "GET",
          headers: setupHeaders()
        })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        return null
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      return null
    })
}

export const deleteUser = (userId) => {
  const url = END_POINT + "/api/1.0/user/" + userId

  return fetch(url, {
          method: "DELETE",
          headers: setupHeaders()
        })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        return null
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      return null
    })
}

export const getUserList = () => {
  const url = END_POINT + "/api/1.0/user/list"

  return fetch(url, {
          method: "GET",
          headers: setupHeaders()
        })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        console.error("Network Error", response)
        return null
      }
    })
    .catch((error) => {
      console.error("Network Error", error)
      return null
    })
}