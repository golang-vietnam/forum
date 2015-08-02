// Webpack config for creating the production bundle.

var path = require('path');
var webpack = require('webpack');
var writeStats = require('./utils/writeStats');
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var strip = require('strip-loader');

var assetsPath = path.join(__dirname, '../static/dist');

module.exports = {
  devtool: 'source-map',
  context: path.resolve(__dirname, '..'),
  entry: {
    'main': './client.js'
  },
  output: {
    path: assetsPath,
    filename: '[name]-[chunkhash].js',
    chunkFilename: '[name]-[chunkhash].js',
    publicPath: '/dist/'
  },
  module: {
    loaders: [
      { test: /\.(jpe?g|png|gif|svg|ttf|eot)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: 'file' },
      { test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "url?limit=10000&minetype=application/font-woff" },
      { test: /\.js$/, exclude: /node_modules/, loaders: [strip.loader('debug'), 'babel?stage=0&optional=runtime&plugins=typecheck']},
      { test: /\.less$/, loader: ExtractTextPlugin.extract('style', 'css!autoprefixer?browsers=last 2 version!less') }
    ]
  },
  progress: true,
  resolve: {
    modulesDirectories: [
      'front',
      'node_modules'
    ],
    extensions: ['', '.json', '.js']
  },
  plugins: [

    // css files from the extract-text-plugin loader
    new ExtractTextPlugin('[name]-[chunkhash].css'),
    new webpack.DefinePlugin({__CLIENT__: true, __SERVER__: false, __DEVELOPMENT__: false, __DEVTOOLS__: false}),

    // ignore dev config
    new webpack.IgnorePlugin(/\.\/dev/, /\/config$/),

    // set global vars
    new webpack.DefinePlugin({
      'process.env': {

        // Mainly used to require CSS files with webpack, which can happen only on browser
        // Used as `if (process.env.BROWSER)...`
        BROWSER: JSON.stringify(true),

        // Useful to reduce the size of client-side libraries, e.g. react
        NODE_ENV: JSON.stringify('production')

      }
    }),

    // optimizations
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
          warnings: false
        }
    }),

    // stats
    function() { this.plugin('done', writeStats); }

  ]
};
