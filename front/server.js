/*global __DEVELOPMENT__*/
import Express from 'express';
import React from 'react';
import Location from 'react-router/lib/Location';
import config from './config';
import favicon from 'serve-favicon';
import compression from 'compression';
import httpProxy from 'http-proxy';
import path from 'path';
import serialize from 'serialize-javascript';
import createStore from './redux/create';
import ApiClient from './ApiClient';
import universalRouter from './universalRouter';
const app = new Express();
const proxy = httpProxy.createProxyServer({
  target: 'http://localhost:' + config.apiPort
});

app.disable('x-powered-by');
app.use(compression());
app.use(favicon(path.join(__dirname, '../favicon.ico')));

let webpackStats;

if (!__DEVELOPMENT__) {
  webpackStats = require('./webpack-stats.json');
}

app.use(require('serve-static')(path.join(__dirname, './static')));

// Proxy to API server
app.use(config.apiPath, (req, res) => {
  proxy.web(req, res);
});

app.use((req, res) => {
  if (__DEVELOPMENT__) {
    webpackStats = require('./webpack-stats.json');
    // Do not cache webpack stats: the script file would change since
    // hot module replacement is enabled in the development env
    delete require.cache[require.resolve('./webpack-stats.json')];
  }
  const client = new ApiClient(req);
  const store = createStore(client);
  const location = new Location(req.path, req.query);
  universalRouter(location, undefined, store)
    .then(({component, transition, isRedirect}) => {
      try {

        if (isRedirect) {
          res.redirect(transition.redirectInfo.pathname);
          return;
        }

        res.send('<!doctype html>\n' + React.renderToString(
            <html lang="en-us">
            <head>
              <meta charSet="utf-8"/>
              <title>Golang Viet Nam</title>
              <link rel="shortcut icon" href="/favicon.ico"/>
              {webpackStats.css.map((css, i) => <link href={css} ref={i}
                                                      media="screen, projection" rel="stylesheet" type="text/css"/>)}
            </head>
            <body>
            <div id="content" dangerouslySetInnerHTML={{__html: React.renderToString(component)}}/>
            <script dangerouslySetInnerHTML={{__html: `window.__data=${serialize(store.getState())};`}}/>
            <script src={webpackStats.script[0]}/>
            </body>
            </html>));
      } catch (error) {
        console.error('ERROR', error.stack);
        res.status(500).send({error: error});
      }
    }, (error) => {
      console.error('ERROR', error.stack);
      res.status(500).send({error: error});
    });
});

if (config.port) {
  app.listen(config.port, (err) => {
    if (err) {
      console.error(err);
    } else {
      console.info('==> Server is listening');
      console.info('==>   %s running on port %s', config.app.name, config.port);
      console.info('==>   %s/ proxy to http://localhost:%s', config.apiPath, config.apiPort);
    }
  });
} else {
  console.error('==> ERROR: No FRONT_PORT environment variable has been specified');
}
