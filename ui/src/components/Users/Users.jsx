import React from 'react'

import { Table, Button } from 'react-bootstrap'
import { NavLink } from 'react-router-dom'
import * as axios from 'axios'

const User = (props) => {
  let path = `/users/${props.state.id}`
  return (
    <tr>
      <td>{props.state.id || "Null"}</td>
      <td><NavLink to={path}>{props.state.login || "Null"}</NavLink></td>
      <td>{props.state.email || "Null"}</td>
      <td>{props.state.firstname || "Null"}</td>
      <td>{props.state.lastname || "Null"}</td>
    </tr>
  );
}

const Users = (props) => {
  let users = props.users.map(v => <User key={v.id} state={v}/>)
  
  const onAddUser = () => {
    props.onAddUser()
  }

  return (
    <>
    <Table striped bordered hover>
       <thead>
        <tr>
          <th>#</th>
          <th>Login</th>
          <th>Email</th>
          <th>First Name</th>
          <th>Last Name</th>
        </tr>
      </thead>
      <tbody>
       {users}
      </tbody>
    </Table>
    <Button onClick={onAddUser}> Новый</Button>
    </> 
  );
}

export default Users