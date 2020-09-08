let initialState = {
  links: [
    {
      to: '/',
      name: 'Главная'
    },
    {
      to: '/users',
      name: 'Сотрудники'
    },
    {
      to: '/units',
      name: 'Торговые точки'
    }
  ]
}

export default (state=initialState, action) => {
  switch (action.type) {
    default:
      return state
  }
}