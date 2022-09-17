import React from 'react'
import { MDBIcon, toast, ToastContainer, MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBRow, MDBCol, MDBInput } from 'mdbreact'
import ToggleButton from '@material-ui/lab/ToggleButton'
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup'
import { FaMale, FaFemale, FaVenusMars, FaBlind, FaBabyCarriage } from 'react-icons/fa'
import Grid from '@material-ui/core/Grid'
import Slider from '@material-ui/core/Slider'
import Image from 'react-bootstrap/Image'
import Female1Img from '../../assets/images/signup/female-1.jpg'
import Female2Img from '../../assets/images/signup/female-2.jpg'
import Female3Img from '../../assets/images/signup/female-3.jpg'
import Female4Img from '../../assets/images/signup/female-4.jpg'
import Female5Img from '../../assets/images/signup/female-5.jpg'
import Female6Img from '../../assets/images/signup/female-6.jpg'
import Female7Img from '../../assets/images/signup/female-7.jpg'
// import Female8Img from '../../assets/images/signup/female-8.jpg'  -- removed as I am mindful of feminism
import Female9Img from '../../assets/images/signup/female-9.jpg'
import Female10Img from '../../assets/images/signup/female-10.jpg'
import Female11Img from '../../assets/images/signup/female-11.jpg'
import Female12Img from '../../assets/images/signup/female-12.jpg'
import Male1Img from '../../assets/images/signup/male-1.jpg'
import Male2Img from '../../assets/images/signup/male-2.jpg'
import Male3Img from '../../assets/images/signup/male-3.jpg'
import Male4Img from '../../assets/images/signup/male-4.jpg'
import Male5Img from '../../assets/images/signup/male-5.jpg'
import Male6Img from '../../assets/images/signup/male-6.jpg'
import Male7Img from '../../assets/images/signup/male-7.jpg'
import Male8Img from '../../assets/images/signup/male-8.jpg'
import Male9Img from '../../assets/images/signup/male-9.jpg'
import Male10Img from '../../assets/images/signup/male-10.jpg'
import Male11Img from '../../assets/images/signup/male-11.jpg'

export default function Welcome({validateRef, signupDetails, setActiveStep}) {
  const submitRef = React.useRef()
  React.useImperativeHandle(validateRef, () => submitRef.current)
  const [name, setName] = React.useState('')
  const [gender, setGender] = React.useState('Male')
  const [age, setAge] = React.useState(7)

  React.useEffect(() => {
    setName(signupDetails.name)
    setGender(signupDetails.gender)
    setAge(signupDetails.age)
  }, [signupDetails])

  const handleGender = (event, newGender) => {
    if (newGender !== null) {
      setGender(newGender)
    }
  }

  const handleAgeChange = (event, newAge) => {
    setAge(newAge)
  }

  function handleSubmit(e) {
    e.preventDefault();

    if (name.length > 0) {
      signupDetails.name = name
      signupDetails.gender = gender
      signupDetails.age = age
      setActiveStep((prevActiveStep) => prevActiveStep + 1)
    } else {
      toast.warn(<span><MDBIcon icon="exclamation-triangle" /> Please provide your Name</span>)
    }
  }

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input type="submit" style={{display: "none"}} ref={submitRef} />
        <MDBRow>
          <MDBCol md='4'>
          </MDBCol>
          <MDBCol md='4'>
            <MDBCard className="text-left">
              <MDBCardBody>
                <MDBCardTitle>Details</MDBCardTitle>
                <MDBCardText>
                  <div className="grey-text">
                    <MDBInput focused label="Full Name (or Nickname)" value={name} onChange={(e) => {setName(e.target.value)} }icon="user" group type="text" />
                    <FaVenusMars className="image-align" color='grey' size={35} />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                    <ToggleButtonGroup
                      value={gender}
                      exclusive
                      onChange={handleGender}>
                      <ToggleButton value="Male">
                        <FaMale color='#4285f4' size={25} />
                      </ToggleButton>
                      <ToggleButton value="Female">
                        <FaFemale color='violet' size={25} />
                      </ToggleButton>
                    </ToggleButtonGroup>
                    <br/>
                    <br/>
                    <div className='font-weight-bold blue-text'>
                    Your Age: {age}
                    </div>
                     <div className='fixed-image-height'>

                        { gender === 'Female' && age >= 1 && age < 3 && <Image src={Female1Img} /> }
                        { gender === 'Female' && age >= 3 && age < 7 && <Image src={Female2Img} /> }
                        { gender === 'Female' && age >= 7 && age < 10 && <Image src={Female3Img} /> }
                        { gender === 'Female' && age >= 10 && age < 14 && <Image src={Female4Img} /> }
                        { gender === 'Female' && age >= 14 && age < 22 && <Image src={Female5Img} /> }
                        { gender === 'Female' && age >= 22 && age < 30 && <Image src={Female6Img} /> }
                        { gender === 'Female' && age >= 30 && age < 42 && <Image src={Female7Img} /> }
                        { gender === 'Female' && age >= 42 && age < 52 && <Image src={Female9Img} /> }
                        { gender === 'Female' && age >= 52 && age < 60 && <Image src={Female10Img} /> }
                        { gender === 'Female' && age >= 60 && age < 85 && <Image src={Female11Img} /> }
                        { gender === 'Female' && age >= 85 && <Image src={Female12Img} /> }

                        { gender === 'Male' && age >= 1 && age < 3 && <Image src={Male1Img} /> }
                        { gender === 'Male' && age >= 3 && age < 7 && <Image src={Male2Img} /> }
                        { gender === 'Male' && age >= 7 && age < 10 && <Image src={Male3Img} /> }
                        { gender === 'Male' && age >= 10 && age < 14 && <Image src={Male4Img} /> }
                        { gender === 'Male' && age >= 14 && age < 22 && <Image src={Male5Img} /> }
                        { gender === 'Male' && age >= 22 && age < 30 && <Image src={Male6Img} /> }
                        { gender === 'Male' && age >= 30 && age < 36 && <Image src={Male7Img} /> }
                        { gender === 'Male' && age >= 36 && age < 44 && <Image src={Male8Img} /> }
                        { gender === 'Male' && age >= 44 && age < 57 && <Image src={Male9Img} /> }
                        { gender === 'Male' && age >= 57 && age < 75 && <Image src={Male10Img} /> }
                        { gender === 'Male' && age >= 75 && <Image src={Male11Img} /> }

                      </div>
                    <Grid container spacing={2}>
                      <Grid item>
                        <FaBabyCarriage color='grey' size={25} />
                      </Grid>
                      <Grid item xs>
                        <Slider min={1} max={100} value={age} onChange={handleAgeChange} aria-labelledby="continuous-slider" />
                      </Grid>
                      <Grid item>
                        <FaBlind color='grey' size={25} />
                      </Grid>
                    </Grid>
                  </div>
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