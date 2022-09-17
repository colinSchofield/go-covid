import React from 'react'
import { MDBIcon, toast, ToastContainer, MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBRow, MDBCol } from 'mdbreact'
import RegionList from './RegionList'

export default function Regions({validateRef, signupDetails, setActiveStep}) {
  const submitRef = React.useRef()
  React.useImperativeHandle(validateRef, () => submitRef.current)
  const [checked, setChecked] = React.useState([])

  React.useEffect(() => {
    setChecked(signupDetails.regions)
  }, [signupDetails])

  function handleSubmit(e) {
    e.preventDefault()

    if (checked.length > 0) {
      signupDetails.regions = checked
      setActiveStep((prevActiveStep) => prevActiveStep + 1)
    } else {
      toast.warn(<span><MDBIcon icon="exclamation-triangle" /> Please select a region</span>)
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
            <MDBCard>
              <MDBCardBody>
                <MDBCardTitle>Regions</MDBCardTitle>
                <MDBCardText className="scrollable">

                  <RegionList checked={checked} setChecked={setChecked} />

                </MDBCardText>
              </MDBCardBody>
            </MDBCard>
          </MDBCol>
          <MDBCol md='4'>
          </MDBCol>
        </MDBRow>
        <br/>
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
