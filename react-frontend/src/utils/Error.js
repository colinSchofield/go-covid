import React from 'react'
import { Alert } from 'react-bootstrap'

export default function Error({ error }) {
  return (
     error &&
      <Alert variant="danger">
      <Alert.Heading><span role="img" aria-labelledby="error">🛑</span> Error</Alert.Heading>
        <p>{error}</p>
      </Alert>
  )
}