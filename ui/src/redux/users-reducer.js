const ADD_USER = 'ADD_USER'
const SET_USERS = 'SET_USERS'

let initialState = {
  users: []
}

export default (state=initialState, action) => {
  switch (action.type) {
    case ADD_USER: {
      return {
        ...state,
        users: [...state.users, {
          id: 5,
          login: 'login',
          email: 'email',
          firstname: 'firstname',
          lastname: 'lastname'
        }]
      }
    }
    case SET_USERS: {
      return {
        ...state,
        users: action.users
      }
    }
    default:
      return state
  }
}