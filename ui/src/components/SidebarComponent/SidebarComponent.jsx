import React from 'react'

import s from './Sidebar.module.css'
import { Nav } from 'react-bootstrap';
import { Link  } from 'react-router-dom'

const SideBarItem = ({href, name, ...props}) => {
  return (
    <Nav.Link as={Link} to={href}>{name}</Nav.Link>
  );
}

const SideBarComponment = (props) => {
  let links = props.links.map(v => <SideBarItem key={v.to} href={v.to} name={v.name}/>)
  return (
    <Nav defaultActiveKey="/" className={`flex-column ${s.sidebar}`}>
      {links}
    </Nav>
  ); 
}

export default SideBarComponment