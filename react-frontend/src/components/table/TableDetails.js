import React from 'react'
import { Line } from 'react-chartjs-2'
import { MDBContainer } from 'mdbreact'
import { Spinner } from 'react-bootstrap'
import { Button, Modal } from 'react-bootstrap'
import { getCovid19Monthly } from '../../utils/api'
import { isMobile } from '../../utils/mobileResponsive'
import Error from '../../utils/Error'
import RegionsContext from '../../context/RegionsContext'

export default function TableDetails({region}) {
  const regionsContext = React.useContext(RegionsContext)
  const [error, setError] = React.useState(null)
  const [tableData, setTableData] = React.useState(null)
  const [show, setShow] = React.useState(false)
  const [text, setText] = React.useState(null)
  const [graphData, setGraphData] = React.useState(null)

  const handleClose = () => {
    setShow(false)
    setText(null)
    setError(null)
    setTableData(null)
    region = null
  }

  const getFlag = (location) => {
    const country = location.trim()
    return regionsContext.regionsData
                  .filter((region) => region.key === country)
                  .map((element) => element.flag + " " + country)
  }

  React.useEffect(() => {
    setText(getFlag(region))
    setShow(true)
    getCovid19Monthly(region)
      .then((stats) => {
        setTableData(stats)
      })
      .catch((exception) => {
        console.log("Error was Caught!", exception)
        setError(exception.message)
      })
  }, [region])


  React.useEffect(() => {
    if (tableData != null) {
      let dataChange = data
      dataChange.labels = tableData.labels
      dataChange.datasets[0].data = tableData.newCases
      dataChange.datasets[1].data = tableData.deaths
      setGraphData(dataChange)

      setText(tableData.flag + " " + region)
      setShow(true)
    }
  }, [tableData, region])

  const data = {
      labels: ["", "", "", "", "", "", ""],
      datasets: [
        {
          label: "New Cases",
          fill: true,
          lineTension: 0.3,
          backgroundColor: "rgba(225, 204,230, .3)",
          borderColor: "rgb(205, 130, 158)",
          borderCapStyle: "butt",
          borderDash: [],
          borderDashOffset: 0.0,
          borderJoinStyle: "miter",
          pointBorderColor: "rgb(205, 130,1 58)",
          pointBackgroundColor: "rgb(255, 255, 255)",
          pointBorderWidth: 10,
          pointHoverRadius: 5,
          pointHoverBackgroundColor: "rgb(0, 0, 0)",
          pointHoverBorderColor: "rgba(220, 220, 220,1)",
          pointHoverBorderWidth: 2,
          pointRadius: 1,
          pointHitRadius: 10,
          data: [0, 0, 0, 0, 0, 0, 0]
        },
        {
          label: "Deaths",
          fill: true,
          lineTension: 0.3,
          backgroundColor: "rgba(184, 185, 210, .3)",
          borderColor: "rgb(35, 26, 136)",
          borderCapStyle: "butt",
          borderDash: [],
          borderDashOffset: 0.0,
          borderJoinStyle: "miter",
          pointBorderColor: "rgb(35, 26, 136)",
          pointBackgroundColor: "rgb(255, 255, 255)",
          pointBorderWidth: 10,
          pointHoverRadius: 5,
          pointHoverBackgroundColor: "rgb(0, 0, 0)",
          pointHoverBorderColor: "rgba(220, 220, 220, 1)",
          pointHoverBorderWidth: 2,
          pointRadius: 1,
          pointHitRadius: 10,
          data: [0, 0, 0, 0, 0, 0, 0]
        }
      ]
    }

  const modalHeight = () => {
    return (isMobile()) ? <p><br/><br/><br/><br/><br/></p> : <p><br/><br/><br/><br/><br/><br/><br/><br/></p>
  }

  return (
    <>
      <Modal show={show} onHide={handleClose} animation={false} size="lg" aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header closeButton>
          <Modal.Title id="contained-modal-title-vcenter"><span className="emoji">{text}</span></Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="App">
          {!tableData && modalHeight() }
          {!error && !tableData && <Spinner animation="border" variant="success" />}
          <Error error={error} />
          {!tableData && modalHeight() }
          </div>

          {tableData && <MDBContainer>
            <h4>New Cases and Deaths over the past month</h4>
            <br/><br/>
            <Line data={graphData} options={{ responsive: true }} />
          </MDBContainer>}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}