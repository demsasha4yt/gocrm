import Units from './Units'
import { connect } from 'react-redux'

let mapStateToProps = (state) => {
  return {
    units: state.unitsPage.units
  }
}

let mapDispatchToProps = (dispatch) => {
  return {}
}

const UnitsContainer = connect(mapStateToProps, mapDispatchToProps)(Units)

export default UnitsContainer