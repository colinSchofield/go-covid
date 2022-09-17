import React from 'react'
import { getUserList } from '../../utils/api'
import { Spinner } from 'react-bootstrap'
import UserTable from './UserTable'
import Error from '../../utils/Error'
import { ToastContainer } from 'mdbreact'

export default function Admin() {
  const [ data, setData ] = React.useState(null)
  const [ error, setError ] = React.useState(null)

  React.useEffect(() => {
    clearTableData()
  }, [])

  const clearTableData = () => {
    setData(null)
    getUserList()
      .then((users) => {
        setData(users)
      })
      .catch((exception) => {
        console.log("Error was Caught!", exception)
        setError(exception.message)
      })
  }

  return (
    <>
      <Error error={error} />

      { !data && !error && <p><br/><br/><br/><br/></p> }
      { !data && !error && <Spinner animation="border" variant="success" /> }

      { data && <UserTable data={data} clearTableData={clearTableData} /> }

      <ToastContainer
        hideProgressBar={true}
        newestOnTop={true}
        autoClose={4000}
      />
    </>
  )
}