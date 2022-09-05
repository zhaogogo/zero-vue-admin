import Vue from 'vue'
import App from './App.vue'

import 'element-ui/lib/theme-chalk/index.css'
import '@/styles/index.scss'
import router from '@/router/index.js'
import store from '@/store/index.js'
import '@/core/core-config.js' 

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')
