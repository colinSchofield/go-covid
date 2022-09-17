import React from 'react'
import { MDBBtn, MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBRow, MDBCol } from 'mdbreact'
import { FaLinkedin, FaEnvelopeSquare, FaComments } from 'react-icons/fa'
import Image from 'react-bootstrap/Image'
import Tooltip from '@material-ui/core/Tooltip'
import MeImg from '../assets/images/me.jpg'
import ReactImg from '../assets/images/react.svg'
import WebPackImg from '../assets/images/web-pack.svg'
import BootStrapImg from '../assets/images/boot-strap.svg'
import FontAwesomeImg from '../assets/images/font-awesome.svg'
import RapidApiImg from '../assets/images/rapid-api.svg'
import AwsImg from '../assets/images/aws.svg'
import GoImg from '../assets/images/Go.svg'
import K8sImg from '../assets/images/K8s.svg'
import GitHubImg from '../assets/images/git-hub.svg'

export default function About() {
  return (
    <MDBRow className='align-middle vertical-align'>
      <MDBCol md='4'>
      </MDBCol>
      <MDBCol md='4'>
        <MDBCard wide cascade>
          <MDBCardBody cascade className='text-center'>
            <Image src={MeImg} roundedCircle width={80} />
            <MDBCardTitle className='card-title'>
              <br/>
              <strong>Colin Schofield</strong>
            </MDBCardTitle>
            <p className='font-weight-bold blue-text'>Senior Go Developer</p>
            <MDBCardText>
              Colin is both a backend Go (golang) and a Full Stack Developer. He has been working in Canada, USA,
              UK &amp; Australia for over fifteen years.
            </MDBCardText>
            <MDBCol>
              <Tooltip title="LinkedIn" arrow>
                <MDBBtn color onClick={() => {window.open(`https://www.linkedin.com/in/colins`, `_self`)}} className="rounded-circle">
                  <FaLinkedin color='#0072b1' size={25}/>
                </MDBBtn>
              </Tooltip>
              <Tooltip title="Mail colin_sch@yahoo.com" arrow>
                <MDBBtn color onClick={() => {window.open(`mailto:colin_sch@yahoo.com`, `_self`)}} className="rounded-circle">
                  <FaEnvelopeSquare color='#808080'size={25} />
                </MDBBtn>
              </Tooltip>
              <Tooltip title="Phone +61 448-644-233" arrow>
                <MDBBtn color onClick={() => {window.open(`sms:+61448644233`, `_self`)}} className="rounded-circle">
                  <FaComments color='#68bd45' size={25}/>
                </MDBBtn>
              </Tooltip>
            </MDBCol>
            <MDBCardText>
              <br/>
              The following tech stack was used to develop this application:
            </MDBCardText>
            <MDBCol md='12'>
              <Tooltip title="Rapid API: CoronaVirus Restful Web Service" arrow>
                <Image src={RapidApiImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="AWS: Online application hosting" arrow>
                <Image src={AwsImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="Webpack: JavaScript module bundler" arrow>
                <Image src={WebPackImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="React: JavaScript library for building user interfaces" arrow>
                <Image src={ReactImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="Font Awesoeme: Font and icon toolkit" arrow>
                <Image src={FontAwesomeImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="Bootstrap: Responsive, mobile-first front-end CSS framework" arrow>
                <Image src={BootStrapImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="Go (golang): written backend microservice" arrow>
                <Image src={GoImg} rounded width='12%' />
              </Tooltip>
              <Tooltip title="Kubernetes (K8s): AWS EKS Managed Kubernetes Cluster" arrow>
                <Image src={K8sImg} rounded width='12%' />
              </Tooltip>
            </MDBCol>
            <MDBCardText>
              <br/>
              <br/>
              <Tooltip title="View source code" placement="right" arrow>
                <a href="https://github.com/colinSchofield/go-covid"><Image className='icon-align-bottom' src={GitHubImg} rounded width={80} /></a>
              </Tooltip>
            </MDBCardText>
          </MDBCardBody>
        </MDBCard>
      </MDBCol>
      <MDBCol md='4'>
      </MDBCol>
    </MDBRow>
  )
}