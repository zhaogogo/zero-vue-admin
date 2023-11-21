import { createApp } from 'vue'
import 'virtual:windi.css'
import './style.css'
import App from './App.vue'
// 引入初始化配置
import core from '@/core/core.js'
import router from '@/router/index.js'

const app = createApp(App)

app.use(core).use(router).mount('#app')

export default app
