import React from 'react'

import {
  Navbar,
  Button
} from 'react-bootstrap'

const NavbarComponent = (props) => {
  return (
    <Navbar variant='dark' bg='dark'>
      <Navbar.Brand href="#home">GoCRM</Navbar.Brand>
      <Navbar.Collapse className="justify-content-end">
        <Button>Войти в систему</Button>
      </Navbar.Collapse>
    </Navbar>
  );
}

export default NavbarComponent