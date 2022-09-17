import React from 'react'
import { MDBIcon, toast, ToastContainer, MDBCard, MDBInput, MDBCardBody, MDBCardTitle, MDBCardText, MDBRow, MDBCol } from 'mdbreact'
import Switch from '@material-ui/core/Switch'
import FormControlLabel from '@material-ui/core/FormControlLabel'

export default function Notification({validateRef, signupDetails, setActiveStep}) {
  const submitRef = React.useRef()
  React.useImperativeHandle(validateRef, () => submitRef.current)
  const [emailChecked, setEmailChecked] = React.useState(false)
  const [smsChecked, setSmsChecked] = React.useState(false)
  const [email, setEmail] = React.useState('')
  const [sms, setSms] = React.useState('')

  React.useEffect(() => {
    if (signupDetails.email !== '' && signupDetails.email !== null ) {
      setEmailChecked(true)
      setEmail(signupDetails.email)
    }
    if (signupDetails.sms !== '' && signupDetails.sms !== null) {
      setSmsChecked(true)
      setSms(signupDetails.sms)
    }
  }, [signupDetails])

  const toggleEmailChecked = () => {
    setEmailChecked((prev) => !prev)
    setEmail('')
  }

  const toggleSmsChecked = () => {
    setSmsChecked((prev) => !prev)
    setSms('')
  }

  function handleSubmit(e) {
    e.preventDefault();
    if (emailChecked && email.length === 0) {
      toast.warn(<span><MDBIcon icon="exclamation-triangle" /> Please provide your email</span>)
    } else if (smsChecked && sms.length === 0) {
      toast.warn(<span><MDBIcon icon="exclamation-triangle" /> Please provide your SMS</span>)
    } else {
      signupDetails.email = email
      signupDetails.sms = sms
      setActiveStep((prevActiveStep) => prevActiveStep + 1)
    }
  }

  return (
    <>
     <form onSubmit={handleSubmit}>
        <input type="submit" style={{display: "none"}} ref={submitRef} />
        <MDBRow className='align-middle'>
          <MDBCol md='4'>
          </MDBCol>
          <MDBCol md='4'>
            <MDBCard className="text-left">
              <MDBCardBody>
                <MDBCardTitle>Notification</MDBCardTitle>
                <MDBCardText>
                  Optionally, provide your email or SMS so we may send you news and updates.<br/><br/>This information shall remain <b>private</b> and will not be shared with anyone.
                  <br/>
                  <br/>
                  <FormControlLabel
                    control={<Switch checked={emailChecked} onChange={toggleEmailChecked} />}
                    label="Notification via Email"
                  />
                  <FormControlLabel
                    control={<Switch checked={smsChecked} onChange={toggleSmsChecked} />}
                    label="Notification via SMS"
                  />

                  { emailChecked && <MDBInput focused value={email} onChange={(e) => { setEmail(e.target.value)}} label="Email Address" icon="envelope" type="email" /> }
                  { smsChecked && <MDBInput focused value={sms} onChange={(e) => { setSms(e.target.value)}} label="SMS Phone number" icon="phone" type="text"  /> }

                </MDBCardText>
              </MDBCardBody>
            </MDBCard>
          </MDBCol>
          <MDBCol md='4'>
          </MDBCol>
        </MDBRow>
        <br/>
      </form>
      <ToastContainer
        hideProgressBar={true}
        newestOnTop={true}
        autoClose={4000}
      />
    </>
  )
}
