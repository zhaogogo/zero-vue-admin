import Vue from 'vue'
import './element_lazy'  // 按需加载element
// 加载网站配置文件夹
import config from './config'
import Bus from '@/utils/bus'
Vue.prototype.$CONFIG = config
Vue.use(Bus)