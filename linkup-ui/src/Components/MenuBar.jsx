import React from 'react';
import {Navbar, Nav, NavDropdown} from 'react-bootstrap';
// import Link from 'react-router-dom/Link';

export class MenuBar extends React.Component {
  render() {
    return (
      <div>
        <Navbar bg="dark" expand="lg" variant="dark">
          <Navbar.Brand href="#">LinkUp</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="mr-auto"></Nav>
          </Navbar.Collapse>
        </Navbar>
      </div>
    );
  }
}

export default MenuBar;
