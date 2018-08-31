const path = require('path');
const webpack = require('webpack');

module.exports = {
  mode: 'development', // 追加
  entry: {  
      app:[path.resolve(__dirname, 'src/components/app.jsx')]
  },
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'assets/js/'),
    publicPath: '/',
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/, 
        use: 'babel-loader',
      },
    ]
  },
  resolve: {
    extensions: ['.js', '.jsx'],
  },
  node: {
    dns: 'mock',
    net: 'mock'
  }
};