import React from 'react'
import Stepper from '@material-ui/core/Stepper'
import Step from '@material-ui/core/Step'
import StepLabel from '@material-ui/core/StepLabel'
import Button from '@material-ui/core/Button'
import Welcome from './Welcome'
import Details from './Details'
import Regions from './Regions'
import Notification from './Notification'
import Confirm from './Confirm'
import RegisterUser from './RegisterUser'
import EditDeleteUser from './EditDeleteUser'
import LoadUser from './LoadUser'
import { getUserIdCookie } from '../../utils/cookies'
import { getUser } from '../../utils/api'

const WELCOME_PAGE = 0
const DETAILS_PAGE = 1
const REGIONS_PAGE = 2
const NOTIFICATION_PAGE = 3
const CONFIRM_PAGE = 4
const REGISTER_USER_PAGE = 5
const EDIT_DELETE_USER = 6
const LOAD_USER = 7

function getSteps() {
  return ['Welcome', 'Details', 'Regions', 'Notification', 'Confirm']
}

export default function SignUp({adminId, returnToAdminTable}) {
  const defaultUserDetails = { id: null, name: "", age: 7, gender: "Male", regions: [], email: "", sms: ""}
  const [details, setDetails] = React.useState(defaultUserDetails)
  const [activeStep, setActiveStep] = React.useState(LOAD_USER)
  const detailsRef = React.useRef()
  const notificationRef = React.useRef()
  const regionsRef = React.useRef()
  const steps = getSteps()

  React.useEffect(() => {
    if (!getUserIdCookie() && !adminId) {
      setActiveStep(WELCOME_PAGE)
      return
    }

    const userId = adminId || getUserIdCookie()
    getUser(userId)
      .then((currentUser) => {
        if (currentUser !== null) {
          setDetails(currentUser)
          setActiveStep(EDIT_DELETE_USER)
        } else {
          setActiveStep(WELCOME_PAGE)
        }
      })
      .catch((exception) => {
        console.log("Error was Caught!", exception)
        setActiveStep(WELCOME_PAGE)
      })
  }, [])

  const handleNext = () => {
    doNextAction(activeStep)
  }

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1)
  }

  function doNextAction(step) {
    switch (step) {
      case WELCOME_PAGE:
        setActiveStep((prevActiveStep) => prevActiveStep + 1)
        break
      case DETAILS_PAGE:
        detailsRef.current.click()
        break
      case REGIONS_PAGE:
        regionsRef.current.click()
        break
      case NOTIFICATION_PAGE:
        notificationRef.current.click()
        break
      case CONFIRM_PAGE:
        setActiveStep((prevActiveStep) => prevActiveStep + 1)
        break
      default:
        throw new Error("Unknown Action")
    }
  }

  function getStepContent(step) {
    switch (step) {
      case WELCOME_PAGE:
        return <Welcome signupDetails={details} setActiveStep={setActiveStep} />
      case DETAILS_PAGE:
        return <Details validateRef={detailsRef} signupDetails={details} setActiveStep={setActiveStep} />
      case REGIONS_PAGE:
        return <Regions validateRef={regionsRef} signupDetails={details} setActiveStep={setActiveStep} />
      case NOTIFICATION_PAGE:
        return <Notification validateRef={notificationRef} signupDetails={details} setActiveStep={setActiveStep} />
      case CONFIRM_PAGE:
        return <Confirm signupDetails={details} />
      case REGISTER_USER_PAGE:
        return <RegisterUser signupDetails={details} returnToAdminTable={returnToAdminTable} />
      case EDIT_DELETE_USER:
        return <EditDeleteUser signupDetails={details} setActiveStep={setActiveStep} returnToAdminTable={returnToAdminTable} />
      case LOAD_USER:
        return <LoadUser />
      default:
        throw new Error("Unknown Step")
    }
}

  return (
    <>
      {  (activeStep < CONFIRM_PAGE) &&
        <Stepper activeStep={activeStep} alternativeLabel>
          {steps.map((label) => (
            <Step key={label}>
              <StepLabel>{label}</StepLabel>
            </Step>
          ))}
        </Stepper>
      }

      <div>
        {  (activeStep >= CONFIRM_PAGE) && <br/> }

        { getStepContent(activeStep) }

        { activeStep < REGISTER_USER_PAGE &&
          <div>
            <div>
              { adminId &&
                <Button variant="contained" onClick={returnToAdminTable}>Cancel</Button>
              }
              <Button
                disabled={activeStep === WELCOME_PAGE}
                onClick={handleBack}
              >
                Back
              </Button>
              <Button variant="contained" color="primary" onClick={handleNext}>
                {activeStep === steps.length - 1 ? 'Finish' : 'Next'}
              </Button>
            </div>
          </div>
        }
      </div>
    </>
  )
}
