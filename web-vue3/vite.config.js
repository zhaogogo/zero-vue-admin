import { defineConfig } from 'vite'
import vuePlugin from '@vitejs/plugin-vue'
import WindiCSS from 'vite-plugin-windicss'



// https://cn.vitejs.dev/config/
export default ({command, mode}) => {
  console.log("command",command)
  console.log("mode",mode)

  const NODE_ENV = mode || "development"
  const envFiles = [
    `.env.${NODE_ENV}`
  ]

  const config = {
      root: "./", //项目根目录（index.html 文件所在的位置）
      base: "./",
      plugins: [vuePlugin(),WindiCSS()],
    }
  return config 
}

