module.exports = {
  development: {
    isProduction: false,
    port: process.env.FRONT_PORT || 9000,
    apiPort: process.env.API_PORT || 8080,
    apiPath: '/v1',
    app: {
      name: 'GolangVN (development)'
    }
  },
  production: {
    isProduction: true,
    port: process.env.FRONT_PORT || 9000,
    apiPort: process.env.API_PORT || 8080,
    apiPath: '/v1',
    app: {
      name: 'GolangVN (production)'
    }
  }
}[process.env.NODE_ENV || 'development'];
