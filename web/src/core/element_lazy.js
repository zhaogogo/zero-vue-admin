/*
* 按需加载element
* */
import Vue from 'vue';
//  按需引入element
import { 
    Form,
    FormItem,
    Input,
    Message,
    MessageBox,
    Button,
    Loading
} from 'element-ui';

Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Button)

Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$message = Message;
Vue.prototype.$loading = Loading.service;
