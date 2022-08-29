const { defineConfig } = require('@vue/cli-service')

const path = require('path')

function resolve(dir) {
  return path.join(__dirname, dir)
}
const name = 'go zero admin vue' 

module.exports = defineConfig({
  publicPath: '/', 
  outputDir: 'dist',
  assetsDir: 'static',
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: false,
  transpileDependencies: true,
  devServer: {
    port: process.env.VUE_APP_CLI_PORT,
    proxy: {
      [process.env.VUE_APP_BASE_API]: {
        target: `${process.env.VUE_APP_BASE_PATH}:${process.env.VUE_APP_SERVER_PORT}`,
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_API]: ''
        }
      }
    }
  },
  configureWebpack: {
    name: name,
    resolve: {
      alias: {
        '@': resolve('src')
      }
    }
  }
})
