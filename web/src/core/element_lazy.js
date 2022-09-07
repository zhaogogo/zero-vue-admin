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
    Loading,
    Container,
    Aside,
    Header,
    Main,
    Menu,
    Submenu,
    MenuItem,
    Col,
    Row,
    Breadcrumb,
    BreadcrumbItem,
    Select,
    Option,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    Badge
} from 'element-ui';
Vue.use(Dropdown)
Vue.use(DropdownItem)
Vue.use(DropdownMenu)
Vue.use(Badge)
Vue.use(Col)
Vue.use(Select)
Vue.use(Option)
Vue.use(Row)
Vue.use(Breadcrumb)
Vue.use(BreadcrumbItem)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Button)
Vue.use(Container)
Vue.use(Aside)
Vue.use(Header)
Vue.use(Main)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)

Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$message = Message;
Vue.prototype.$loading = Loading.service;
