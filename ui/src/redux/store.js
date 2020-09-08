import {createStore, combineReducers} from 'redux'
import usersReducer from './users-reducer';
import sidebarReducer from './sidebar-reducer';
import unitsReducer from './units-reducer';

let reducers = combineReducers({
  usersPage: usersReducer,
  unitsPage: unitsReducer,
  sidebar: sidebarReducer
})

let store = createStore(reducers);

export default store