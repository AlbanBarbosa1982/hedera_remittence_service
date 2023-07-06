const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/api', // You can setup a specific endpoint such as /api, or use '*' to forward all requests
    createProxyMiddleware({
      target: 'http://localhost/api/transfer:8000', // The address of your backend server
      changeOrigin: true,
    })
  );
};
