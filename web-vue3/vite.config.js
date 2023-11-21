import vuePlugin from '@vitejs/plugin-vue'
import { loadEnv } from 'vite'
import WindiCSS from 'vite-plugin-windicss'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import * as path from 'path'


// https://cn.vitejs.dev/config/
/** @type {import('vite').UserConfig} */
export default ({command, mode}) => {
  console.log("command",command)
  console.log("mode",mode)
  console.log("__dirname",__dirname)
  // console.log("process", process) // process是node的进程启动时的环境变量

  // 当调用loadEnv的时候
  //  1. 直接找到.env文件，解析其中的环境变量 XXX = a
  //  2. 会将传进来的变量mode的值和 .env. 进行拼接，并根据我们提供的目录[loadEnv函数的第二个参数]，解析文件中的k = v值作为环境变量
  // loadEnv参数
  //    根据当前工作目录中的 `mode` 加载 .env 文件
  //    设置第三个参数为 '' 来加载所有环境变量，而不管是否有 `VITE_` 前缀。
  //  而客户端使用环境变量时，import.meta.env 变量中 
  //    需要VITE_开头的环境变量才会注入
  //    通过envPrefix来配置
  const env = loadEnv(mode, process.cwd(),'')

  const config = {
      root: "./", //项目根目录（index.html 文件所在的位置）
      base: "./", // js导入的资源路径，src
      resolve: {
        alias: {
          "@": path.resolve(__dirname, "./src")
        }
      },
      plugins: [
        vuePlugin(),
        WindiCSS(),
        AutoImport({
          resolvers: [ElementPlusResolver()],
          // 自动引入 import { reactive, ref } from 'vue'  
          imports: ["vue","vue-router"]
        }),
        Components({
          resolvers: [ElementPlusResolver()],
        }),
      ],
      envPrefix: "VITE_",
    }
  return config 
}

