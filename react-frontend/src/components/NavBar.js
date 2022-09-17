import React from 'react'
import { Navbar, Nav } from 'react-bootstrap'
import TableSearch from './table/TableSearch'

export default function NavBar({ showSearch }) {
  return (
    <Navbar bg="light" expand="lg" variant="light">
      <Navbar.Brand href="/"><span role="img" aria-labelledby="virus" className="emoji">🦠</span> Covid-19 Daily Statistics <span role="img" aria-labelledby="virus" className="emoji">🦠</span></Navbar.Brand>
      <Navbar.Toggle />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="mr-auto">
        <Nav.Link href="/">Home</Nav.Link>

          { /* Note: the html extension -- this is to avoid requiring SSR with Tomcat (i.e. each file is a COPY of the minified index.html file) */ }

          <Nav.Link href="/about.html">About</Nav.Link>
          <Nav.Link href="/signup.html">Sign Up</Nav.Link>
          <Nav.Link href="/admin.html">Admin</Nav.Link>
        </Nav>
        { showSearch && <TableSearch /> }
      </Navbar.Collapse>
    </Navbar>
  )
}