import React from 'react'
import { getCovid19Daily } from '../utils/api'
import { Spinner } from 'react-bootstrap'
import Table from './table/Table'
import DataContext from '../context/DataContext'
import Error from '../utils/Error'
import { ToastContainer } from 'mdbreact'
import { getUser } from '../utils/api'
import { getUserIdCookie, setSessionCookie, isSessionCookieSet } from '../utils/cookies'
import { useCountryFromGeoLocation } from '../utils/country'

export default function Home() {
  const [ data, setData ] = React.useState(null)
  const [ region, setRegion ] = React.useState(null)
  const [ error, setError ] = React.useState(null)
  const [ highlightRegions, setHighlightRegions] = React.useState(null)
  const tableContext = React.useContext(DataContext)
  const [ country, setCountry ] = React.useState(null)
  useCountryFromGeoLocation(setCountry)

  React.useEffect(() => {
    getCovid19Daily()
      .then((stats) => {
        setData(stats.response)
        tableContext.updateTable({data: stats.response, countrySelected: null})
      })
      .catch((exception) => {
        console.log("Error was Caught!", exception)
        setError(exception.message)
      })
  }, [])

  React.useEffect(() => {
    if (isSessionCookieSet()) {
      return
    }
    if (country) {
      setSessionCookie()    // Only display the geo-located graph once per session
      setRegion(country)
      if (!getUserIdCookie()) {
        var highlight = {region: [ country ]}
        setHighlightRegions(highlight.region)
      }
    }
  }, [country])

  React.useEffect(() => {
    if (!getUserIdCookie()) {
      return
    }
    getUser(getUserIdCookie())
      .then((currentUser) => {
        if (currentUser !== null) {
          setHighlightRegions(currentUser.regions)
        }
      })
  }, [])

  React.useEffect(() => {

    if (tableContext.tableData !== null &&
        tableContext.tableData.countrySelected !== null) {

      setRegion(tableContext.tableData.countrySelected)
    } else if (tableContext.tableData !== null &&
               tableContext.tableData.data !== null) {

      setData(tableContext.tableData.data)
    }
  }, [tableContext])

  return (

    <>
      <Error error={error} />

      { !data && !error && <p><br/><br/><br/><br/></p> }
      { !data && !error && <Spinner animation="border" variant="success" /> }

      { data && <Table data={data} displayRegion={region} highlightRegions={highlightRegions} /> }

      <ToastContainer
        hideProgressBar={true}
        newestOnTop={true}
        autoClose={4000}
      />
    </>

  )
}