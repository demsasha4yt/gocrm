import React from 'react'

import './App.css'

import {Container, Row, Col} from 'react-bootstrap'
import {
  Route
} from "react-router-dom"

import NavbarComponent from './components/NavbarComponent';

import SideBarComponmentContainer from './components/SidebarComponent';
import UsersContainer from './components/Users'
import UnitsContainer from './components/Units';
import Home from './components/Home/index';

function App(props) {
  return (
    <>
      <NavbarComponent />
        <Container fluid>
          <Row noGutters className='fluid app-wrapper'>
            <Col md={2} className="no-float">
              <SideBarComponmentContainer />
            </Col>
            <Col md={10} className="no-float app-wrapper-content">
              <Route exact path='/' render={() => <Home />} />
              <Route exact path='/users' render={() => <UsersContainer />} />
              <Route exact path='/units' render={() => <UnitsContainer />} />
            </Col>
          </Row> 
      </Container>
    </> 
  );
}

export default App  