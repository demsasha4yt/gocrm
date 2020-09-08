const ADD_UNIT = 'ADD_UNIT'

let initialState = {
  units: [
    {id: 1},
    {id: 2},
  ]
}

export default (state=initialState, action) => {
  switch (action.type) {
    case ADD_UNIT:
      let newUnit = {
        id: 5,
      }
      state.units.push(newUnit)
      return state
    default:
      return state
  }
}