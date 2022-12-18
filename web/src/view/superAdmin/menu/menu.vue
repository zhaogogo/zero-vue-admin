<template>
    <div>
        <div>
            <el-button type="primary" icon="el-icon-plus" size="mini" style="float:right" @click="addMenu(0)">添加根菜单</el-button>
        </div>
        <el-table :data="tableData" border row-key="id" stripe>
            <el-table-column label="id" min-width="100" prop="id"></el-table-column>
            <el-table-column label="展示名称" min-width="120">
                <template slot-scope="scop">
                    <span>{{scop.row.meta.title}}</span>
                </template>
            </el-table-column>
            <el-table-column label="图标" min-width="140">
                <template slot-scope="scope">
                    <i :class="`el-icon-${scope.row.meta.icon}`"></i>
                    <span>{{scope.row.meta.icon}}</span>
                </template>
            </el-table-column>
            <el-table-column label="路由Name" min-width="100" prop="name"></el-table-column>
            <el-table-column label="路由Path" min-width="100" prop="path"></el-table-column>
            <el-table-column label="是否隐藏" min-width="100">
                <template slot-scope="scope">
                    <span>{{scope.row.hidden?"隐藏":"显示"}}</span>
                </template>
            </el-table-column>
            <el-table-column label="父节点" min-width="90" prop="parentId"></el-table-column>
            <el-table-column label="排序" min-width="70" prop="sort"></el-table-column>
            <el-table-column label="文件路径" min-width="360" prop="component"></el-table-column>
            <el-table-column fixed="right" label="操作" width="300">
                <template slot-scope="scope">
                    <el-button size="mini" type="primary" icon="el-icon-edit" @click="addMenu(scope.row.id)">添加子菜单</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-edit" @click="editMenu(scope.row)">编辑</el-button>
                    <el-button size="mini" type="danger" icon="el-icon-delete" @click="deleteMenu(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        
        <el-dialog :before-close="handleClose" :title="dialogTitle" :visible.sync="dialogFormVisible">
            <el-form ref="menuForm" :inline="true" :model="form" :rules="rules" label-position="top" label-width="70px">
                <el-form-item label="路由name" prop="path" style="width:30%">
                    <el-input v-model="form.name" autocomplete="off" placeholder="唯一英文字符串" @change="changeName"></el-input>
                </el-form-item>
                <el-form-item prop="path" style="width: 30%">
                    <div slot="label" style="display:inline">
                        路由path
                        <el-checkbox v-model="checkFlag" style="display:flex;float:right;margin-left:10px;margin-top:10px">添加参数</el-checkbox>
                    </div>
                    <el-input v-model="form.path" :disabled="!checkFlag" autocomplete="off" placeholder="建议只在后方拼接参数"></el-input>
                </el-form-item>
                <el-form-item label="是否隐藏" style="width:30%">
                    <el-select v-model="form.hidden" placeholder="是否在列表隐藏">
                        <el-option :value="false" label="否"></el-option>
                        <el-option :value="true" label="是"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="父节点Id" style="width: 30%">
                    <el-cascader 
                        v-model.number="form.parentId" 
                        :disabled="!isEdit" 
                        :options="menuOption"
                        :props="{checkStrictly:true,label:'title',value:'id',disabled:'disabled',emitPath:false}"
                        :show-all-levels="false"
                        filterable
                    ></el-cascader>
                </el-form-item>
                <el-form-item label="文件路径" prop="component" style="width:60%">
                    <el-input v-model="form.component" autocomplete="off"></el-input>
                    <span style="font-size: 12px;margin-right:12px">如果菜单包含子菜单，请创建router-view二级路由页面或者</span><el-button size="mini" @click="form.component = 'view/routerHolder.vue'">点我设置</el-button>
                </el-form-item>
                <el-form-item label="页面title" prop="title" style="width:30%">
                    <el-input v-model="form.title" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="图标" prop="icon" style="width:30%">
                    <icon :meta="form.meta">
                        <template slot="prepend">el-icon-</template>
                    </icon>
                </el-form-item>
                <el-form-item label="排序标记" prop="sort" style="width:30%">
                    <el-input v-model.number="form.sort" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>

            <div v-if="isEdit">
                <div style="color:#dc143c">新增菜单需要在角色管理内配置权限才可使用</div>
                <el-button
                    size="small"
                    type="primary"
                    icon="el-icon-edit"
                    @click="addParameter(form)"
                >新增菜单参数</el-button>
                <el-table :data="form.parameters" stripe style="width:100%">
                    <el-table-column prop="type" label="参数类型" width="100">
                        <template slot-scope="scope">
                            <el-select v-model="scope.row.type" placeholder="请选择">
                                <el-option key="query" value="query" label="query"></el-option>
                                <el-option key="params" value="params" label="params"></el-option>
                            </el-select>
                        </template>
                    </el-table-column>
                    <el-table-column prop="user_id" label="用户参数" width="150">
                        <template slot-scope="scope">
                            <el-cascader
                                v-model="scope.row.user_id"
                                :options="userOption"
                                :props="{checkStrictly: true, label:'username', value:'userid', disabled:'disabled', emitPath: false}"
                                clearable
                                filterable
                            >

                            </el-cascader>
                        </template>
                    </el-table-column>
                    <el-table-column prop="key" label="参数key" width="100">
                        <template slot-scope="scope">
                            <el-input v-model="scope.row.key"></el-input>
                        </template>
                    </el-table-column>
                    <el-table-column prop="value" label="参数值">
                        <template slot-scope="scope">
                            <el-input v-model="scope.row.value"></el-input>
                        </template>
                    </el-table-column>
                    <el-table-column>
                        <template slot-scope="scope">
                            <el-button type="danger" size="small" icon="el-icon-delete" @click="deleteParameter(form.parameters,scope.$index)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div slot="footer">
                <el-button @click="closeDialog">取消</el-button>
                <el-button type="primary" @click="enterDialog">确定</el-button>
            </div>
        </el-dialog> 
    </div>
</template>

<script>

import { 
    allMenu,
    addMenu
} from "@/api/menu/menu"

import { getAllUser } from '@/api/user/user'

import infoList from "@/mixins/infoList"
import icon from '@/view/superAdmin/menu/icon.vue'

export default {
    name:"Menu",
    components:{icon},
    data(){
        return {
            checkFlag: false,
            listApi: allMenu,
            dialogFormVisible: false,
            dialogTitle: "",
            menuOption: [
                {
                    id: 0,
                    title:"根菜单"
                }
            ],
            userOption:[],
            form: {
                path: '',
                name: '',
                hidden: false,
                parentId: '',
                component: '',
                title: '',
                icon: '',
                parameters: []
            },
            rules: {
                parentId: [
                    {required: true, message:"输入菜单名字", trigger: 'blur'},
                    {type: "number", message:"必须是数字", trigger: 'blur'}
                ],
                name: [
                    {required: true, message:"输入菜单名字", trigger: 'blur'}
                ],
                path: [
                    {required: true, message:"输入菜单路径", trigger: 'blur'}
                ],
                component: [
                    {required:true, message:"输入文件路径", trigger:'blur'}
                ],
                title: [
                    {required: true, message:"输入菜单展示名称", trigger:'blur'}
                ]
            },
            isEdit: false,
            test:''
        }
    },
    mixins:[infoList],
    async created(){
        this.getTableData()
        const alluser = await getAllUser()
        if (alluser.code === 200) {
            this.setUserOptions(alluser.list, this.userOption)
        }
    },
    methods:{
        // 添加菜单方法，id为 0则为添加根菜单
        addMenu(id){
            this.dialogTitle = "新增菜单"
            this.form.parentId = id
            this.isEdit = false
            this.setOptions()
            this.dialogFormVisible = true
        },
        //父节点Id设置
        setOptions(){
            this.menuOption = [
                {
                    'id': 0,
                    'title':'根目录'
                }
            ]
            this.setMenuOptions(this.tableData, this.menuOption,false)
        },
        setMenuOptions(menuData,optionsData,disabled){
            menuData && menuData.map(item => {
                if(item.children && item.children.length){
                    const option = {
                        title: item.meta.title,
                        id: item.id,
                        disabled: disabled || item.id === this.form.id,
                        children: []
                    }
                    this.setMenuOptions(item.children, option.children,disabled || item.id === this.form.id)
                    optionsData.push(option)
                }else {
                    const option = {
                        title: item.meta.title,
                        id: item.id,
                        disabled: disabled || item.id === this.form.id
                    }
                    optionsData.push(option)
                }
            })
        },
        setUserOptions(userdata, optionsdata){
            userdata && userdata.map(item => {
                const option = {
                    userid: item.id,
                    username: item.name
                }
                optionsdata.push(option)
            })
        },
        //新增用户菜单参数 前端设置
        addParameter(form){
            console.log(form)
            if (!form.parameters){
                this.$set(form,'parameters',[])
            }
            form.parameters.push({
                type:"query",
                key:'',
                value:''
            })
        },
        deleteParameter(parameters, index){
            parameters.splice(index,1)
        },
        changeName(){
            this.form.path = this.form.name
        },
        handleClose(done){
            this.initForm()
            done()
        },
        // deleteMenu(row){
        //     this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        //         confirmButtonText: '确定',
        //         cancelButtonText: '取消',
        //         type: 'warning'
        //     })
        //         .then(async () => {
        //             const res = await deleteBaseMenu({id: row.id})
        //             if (res.code === 0){
        //                 this.$message({
        //                     type:"success",
        //                     message:res.msg
        //                 })
        //                 if (this.tableData.length === 1 && this.page > 1) {
        //                     this.page--
        //                 }
        //                 this.getTableData()
        //             }
        //         })
        //         .catch(() => {
        //             this.$message({
        //                 type:"info",
        //                 message:"已取消删除"
        //             })
        //         })
        // },
        initForm(){
            this.checkFlag = false
            this.$refs.menuForm.resetFields()
            this.form = {
                id: 0,
                path: '',
                name: '',
                hidden: false,
                parentId: '',
                component: '',
                meta:{
                    title: '',
                    icon: '',
                },
                parameters: []
            }
        },
        closeDialog(){
            this.initForm()
            this.dialogFormVisible = false
        },
        enterDialog(){
            this.$refs.menuForm.validate(async(valid)=>{
                if(valid) {

                    let res 
                    if (this.isEdit) {
                        res = await updateBaseMenu(this.form)
                    }else {
                        res = await     addMenu(this.form)
                    }
                    if (res.code === 200) {
                        this.$message({
                            type: 'success',
                            message: res.msg
                        })
                        this.getTableData()
                        this.initForm()
                        this.dialogFormVisible = false
                    }
                }
            })
        },
        async editMenu(row) {
            this.dialogTitle = '编辑菜单'
            // const res = await getBaseMenuById({id})
            this.form = JSON.parse(JSON.stringify(row))
            this.isEdit = true
            this.setOptions()
            this.dialogFormVisible = true
        }
    }
}
</script>

<style>

</style>