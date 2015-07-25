import React from 'react';
import {Route} from 'react-router';
import App from 'views/App';
import Home from 'views/Home';
import About from 'views/About';
import NotFound from 'views/NotFound';
import Redirect from 'views/Redirect';

export default (
  <Route component={App}>
    <Route path="/" component={Home}/>
    <Route path="*" component={NotFound}/>
  </Route>
);
