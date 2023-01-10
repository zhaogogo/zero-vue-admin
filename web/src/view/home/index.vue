<template>
    <div class="user-addcount">
        <el-tabs @tab-click="handleClick">
            <el-tab-pane label="账号绑定" name="second">
                <ul>
                    <li>
                        <p class="title">修改密码</p>
                        <p class="desc">
                        修改个人密码
                        <a href="#" @click="showPassword=true">修改密码</a>
                        </p>
                    </li>
                </ul>
            </el-tab-pane>
        </el-tabs>

        <el-dialog :visible.sync="showPassword" title="修改密码" width="360px" @close="clearPassword">
        <el-form ref="modifyPwdForm" :model="pwdModify" :rules="rules" label-width="80px">
            <el-form-item :minlength="6" label="新密码" prop="password">
            <el-input v-model="pwdModify.password" show-password />
            </el-form-item>
            <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
            <el-input v-model="pwdModify.confirmPassword" show-password />
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="showPassword=false">取 消</el-button>
            <el-button type="primary" @click="savePassword">确 定</el-button>
        </div>
        </el-dialog>
    </div>
</template>

<script>
import {changeLoginPassword} from '@/api/system/user/user'
export default {
    name: "PersionHome",
    data() {
        return {
            showPassword: false,
            pwdModify: {},
            rules: {
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, message: '最少6个字符', trigger: 'blur' }
                ],
                confirmPassword: [
                    { required: true, message: '请输入确认密码', trigger: 'blur' },
                    { min: 6, message: '最少6个字符', trigger: 'blur' },
                    {
                        validator: (rule, value, callback) => {
                        if (value !== this.pwdModify.password) {
                            callback(new Error('两次密码不一致'))
                        } else {
                            callback()
                        }
                        },
                        trigger: 'blur'
                    }
                ]
            }
        }
    },
    methods: {
        handleClick(tab, event) {
            console.log("------>",tab, event)
        },
        clearPassword() {
            this.pwdModify = {
                password: '',
                confirmPassword: ''
            }
            this.$refs.modifyPwdForm.clearValidate()
        },
        savePassword() {
            this.$refs.modifyPwdForm.validate(valid => {
                if (valid) {
                    changeLoginPassword({
                        password: this.pwdModify.password
                    })
                    .then((res) => {
                        if (res.code === 200) {
                            this.$message.success('修改密码成功！')
                        }
                        this.showPassword = false
                    })
                } else {
                        return false
                    }
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }
      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;
        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }
      border-bottom: 2px solid #f0f2f5;
    }
  }
}
</style>