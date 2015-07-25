import React, {Component, PropTypes} from 'react';
import {Link} from 'react-router';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import {isLoaded as isInfoLoaded} from '../reducers/info';
import {createTransitionHook} from '../universalRouter';
import Header from '../components/layout/Header';

class App extends Component {
  static propTypes = {
    logout: PropTypes.func
  }

  static contextTypes = {
    router: PropTypes.object.isRequired,
    store: PropTypes.object.isRequired
  };

  componentWillMount() {
    const {router, store} = this.context;
    router.addTransitionHook(createTransitionHook(store));
  }

  handleLogout(event) {
    event.preventDefault();
    this.props.logout();
  }

  render() {
    return (
      <div className="app container">
        <Header/>
        {this.props.children}
      </div>
    );
  }
}

@connect(state => ({

}))
export default
class AppContainer {
  static propTypes = {
    dispatch: PropTypes.func.isRequired
  }

  static fetchData(store) {
    const promises = [];
    return Promise.all(promises);
  }

  render() {
    const { dispatch } = this.props;
    return <App>
      {this.props.children}
    </App>;
  }
}
