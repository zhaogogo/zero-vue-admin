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
    Badge,
    Scrollbar,
    Divider,
    Drawer,
    ColorPicker,
    Tabs,
    TabPane,
    Table,
    TableColumn,
    Cascader,
    Dialog,
    Checkbox,
    Tag,
    Switch,
    Tooltip,
    Pagination,
    Popover
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
Vue.use(Scrollbar)
Vue.use(Divider)
Vue.use(Drawer)
Vue.use(ColorPicker)
Vue.use(Tabs)
Vue.use(TabPane)
Vue.use(Loading.directive)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Cascader)
Vue.use(Dialog)
Vue.use(Checkbox)
Vue.use(Tag)
Vue.use(Switch)
Vue.use(Tooltip)
Vue.use(Pagination)
Vue.use(Popover)

Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$message = Message;
Vue.prototype.$loading = Loading.service;
