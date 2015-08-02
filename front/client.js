/* global __DEVTOOLS__ */
require('./styles/main.less');

import React from 'react';
import BrowserHistory from 'react-router/lib/BrowserHistory';
import Location from 'react-router/lib/Location';
import createStore from './redux/create';
import ApiClient from './ApiClient';
import universalRouter from './universalRouter';

const history = new BrowserHistory();
const client = new ApiClient();

const dest = document.getElementById('content');
const store = createStore(client, window.__data);
const location = new Location(document.location.pathname, document.location.search);
universalRouter(location, history, store)
  .then(({component}) => {
    if (__DEVTOOLS__) {
      window.__store = store;

      let state = JSON.stringify(store.getState());
      let states = window.__states = [JSON.parse(state)];
      console.log('state', state);

      store.subscribe(() => {
        let state = JSON.stringify(store.getState());
        states.push(JSON.parse(state));
        console.log('state', state);
      });
    }

    React.render(component, dest);
  }, (error) => {
    console.error(error);
  });


if (process.env.NODE_ENV !== 'production') {
  window.React = React; // enable debugger
  const reactRoot = window.document.getElementById('content');

  if (!reactRoot || !reactRoot.firstChild || !reactRoot.firstChild.attributes || !reactRoot.firstChild.attributes['data-react-checksum']) {
    console.error('Server-side React render was discarded. Make sure that your initial render does not contain any client-side code.');
  }
}
