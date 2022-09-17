import React from 'react'
import { MDBIcon, toast } from 'mdbreact'
import { useHistory } from 'react-router-dom'
import { createUser, updateUser } from '../../utils/api'
import { setUserIdCookie } from '../../utils/cookies'
import { Spinner } from 'react-bootstrap'
import Error from '../../utils/Error'

export default function RegisterUser({signupDetails, returnToAdminTable}) {
  const history = useHistory()
  const [ error, setError ] = React.useState(null)
  const isAdmin = window.location.href.endsWith('/admin.html')

  React.useEffect(() => {

    if (signupDetails.email === null || signupDetails.email.length === 0) {
      signupDetails.email = null
    }
    if (signupDetails.sms === null || signupDetails.sms.length === 0) {
      signupDetails.sms = null
    }

    if (signupDetails.id === null) {
      createUser(signupDetails)
        .then((response) => {
          setUserIdCookie(response.id)
          window.setTimeout(() => toast.success(<span><MDBIcon far icon="check-circle" /> Thanks for registering {signupDetails.name}! </span>), 800)
          history.push("/")

        })
        .catch((exception) => {
          console.log("Error was Caught!", exception)
          setError(exception.message)
        })
      } else {
        updateUser(signupDetails)
          .then((response) => {

            if (isAdmin) {
              window.setTimeout(() => toast.success(<span><MDBIcon far icon="check-circle" /> Details were updated for {signupDetails.name}. </span>), 800)
              returnToAdminTable()
            } else {
              setUserIdCookie(response.id)
              window.setTimeout(() => toast.success(<span><MDBIcon far icon="check-circle" /> Details were updated {signupDetails.name}! </span>), 800)
              history.push("/")
            }
          })
          .catch((exception) => {
            console.log("Error was Caught!", exception)
            setError(exception.message)
          })
        }
  }, [])

  return (
    <>
      <Error error={error} />
      { !error && <p><br/><br/><br/></p> }
      { !error && <Spinner animation="border" variant="success" /> }
    </>
  )
}