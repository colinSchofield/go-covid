import React from 'react'
import { MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBRow, MDBCol } from 'mdbreact'
import Image from 'react-bootstrap/Image'
import BannerImg from '../../assets/images/Coronavirus-Banner5.jpg'
import MapLayout from './MapLayout'

export default function Confirm({signupDetails}) {
  return (
    <>
      <MDBRow className='align-middle'>
        <MDBCol md='4'>
        </MDBCol>
        <MDBCol md='4'>
          <MDBCard>
            <Image className="img-fluid" src={BannerImg} />
            <MDBCardBody>
              <MDBCardTitle>Confirmation</MDBCardTitle>
              <MDBCardText>
              Click <b>Finish</b> to save your details:
              <br/>
              <br/>
              <table className="table table-striped table-hover table-bordered table-sm text-left">
                <tr>
                  <td className='font-weight-bold blue-text'>Name</td>
                  <td>{signupDetails.name}</td>
                </tr>
                <tr>
                  <td className='font-weight-bold blue-text'>Age</td>
                  <td>{signupDetails.age}</td>
                </tr>
                <tr>
                  <td className='font-weight-bold blue-text'>Identify as</td>
                  <td>{signupDetails.gender}</td>
                </tr>

                { signupDetails.email !== '' &&
                  <tr>
                    <td className='font-weight-bold blue-text'>Notification via Email</td>
                    <td>{signupDetails.email}</td>
                  </tr>
                }

                { signupDetails.sms !== '' &&
                  <tr>
                    <td className='font-weight-bold blue-text'>Notification via SMS</td>
                    <td>{signupDetails.sms}</td>
                  </tr>
                }
                <tr>
                  <td className='font-weight-bold blue-text'>Regions</td>
                  <td>
                    <MapLayout regions={signupDetails.regions} />
                  </td>
                </tr>
              </table>
              </MDBCardText>
            </MDBCardBody>
          </MDBCard>
        </MDBCol>
        <MDBCol md='4'>
        </MDBCol>
      </MDBRow>
      <br/>
    </>
  )
}