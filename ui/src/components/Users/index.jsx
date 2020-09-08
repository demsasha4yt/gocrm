import Users from './Users';
import { connect } from 'react-redux';

let mapStateToProps = (state) => {
  return {
    users: state.usersPage.users  
  }
}

let mapDispatchToProps = (dispatch) => {
  return {
    onAddUser: () => {
      dispatch({type: "ADD_USER"})
    },
    setUsers: (users) => {
      dispatch({type: "SET_USERS", users})
    }
  }
}

const UsersContainer = connect(mapStateToProps, mapDispatchToProps)(Users)

export default UsersContainer