import React from 'react'
import ReactDOM from 'react-dom'
import './css/index.css'
import App from './App'

// Import the Bootstrap Material Design CSS files
import "@fortawesome/fontawesome-free/css/all.min.css"
import "mdbreact/dist/css/mdb.css"

// Importing the Bootstrap CSS files
import 'bootstrap/dist/css/bootstrap.min.css'
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css'

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
)