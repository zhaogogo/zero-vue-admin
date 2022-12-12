<template>
    <div class="login-container">
        <el-form 
            ref="loginForm" 
            :model="loginForm" 
            :rules="rules" 
            class="login-form" 
            auto-complete="on" 
            label-position="left"
            @keyup.enter.native="loginHandle"
        >
            <div class="login-title-container">
                <h3 class="login-title">登陆</h3>
            </div>
            
            <el-form-item prop="userName">
                <el-input v-model="loginForm.userName" placeholder="请输入用户名">
                    <i slot="suffix" class="el-input__icon el-icon-user"></i>
                </el-input>
            </el-form-item>
            <el-form-item prop="passWord">
                <el-input v-model="loginForm.passWord" placeholder="请输密码" :type="lock === 'lock' ? 'password' : 'text'">
                    <i class="show-pwd" slot="suffix" :class="'el-input__icon el-icon-' + lock" @click="changeLock"></i>
                </el-input>
            </el-form-item>
            
            <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px" @click="loginHandle">登陆</el-button>
        </el-form>
    </div>
</template>

<script>
import { mapActions } from 'vuex'


export default {
    name: "Login",
    data(){
        const validateUserName = (rule, value, callback) => {
            if (value.length < 3) {
                return callback(new Error('请输入正确的用户名'))
            }else {
                callback()
            }
        }
        const validatePassWord = (rule, value, callback) => {
            if(value.length < 2) {
                return callback(new Error('请输入正确的密码'))
            }else {
                callback()
            }
        }
        return {
            loginForm: {
                userName: "admin",
                passWord: "123"
            },
            lock: 'lock',
            loading: false,
            rules: {
                userName: [{validator: validateUserName, trigger: 'blur'}],
                passWord: [{validator: validatePassWord, trigger: 'blur'}]
            }
        }
    },
    methods:{
        ...mapActions('user',['loginin']),
        loginHandle(){
            this.$refs.loginForm.validate(async (v)=> {
                if (v) {
                    this.loading = true
                    await this.loginin(this.loginForm)
                    this.loading = false
                }
            })
        },
        changeLock() {
            this.lock = this.lock === "lock" ? "unlock" : "lock"
        }
    }
}
</script>

<style lang="scss">

$bg:#283443;
$light_gray:#fff;
$cursor: #fff;
/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    height: 100%;
    width: 100%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>

<style lang="scss" scoped>
    $bg:#2d3a4b;
    $dark_gray:#889aa4;
    $light_gray:#eee;
    
    .login-container {
      min-height: 100%;
      width: 100%;
      background-color: $bg;
      overflow: hidden;
    
      .login-form {
        position: relative;
        width: 520px;
        max-width: 100%;
        padding: 160px 35px 0;
        margin: 0 auto;
        overflow: hidden;
      }
    
      .login-title-container {
        position: relative;
    
        .login-title {
          font-size: 26px;
          color: $light_gray;
          margin: 0px auto 40px auto;
          text-align: center;
          font-weight: bold;
        }
      }
    
      .show-pwd {
        position: static;
        right: 10px;
        top: 7px;
        font-size: 16px;
        color: $dark_gray;
        cursor: pointer;
        user-select: none;
      }
    }
</style>