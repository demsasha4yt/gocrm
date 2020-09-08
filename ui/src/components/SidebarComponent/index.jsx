import { connect } from 'react-redux'
import SideBarComponment from './SidebarComponent';

let mapStateToProps = (state) => {
  return {
    links: state.sidebar.links
  }
}

let mapDispatchToProps = (dispatch) => {
  return {}
}

const SideBarComponmentContainer = connect(mapStateToProps, mapDispatchToProps)(SideBarComponment)

export default SideBarComponmentContainer