<template>
  <div 
    class="h-full w-full grid"
    style="background-color: #2d3a4b;"
  >
  <!-- w-3/5 h-1/5 absolute inset-48 rounded-4xl -->
    <div class="place-self-center w-1/3">
      <el-form 
        class="p-4"
        ref="loginForm"
        :model="loginFromData"
        :rules="loginFromRules"
      >
        <h3 class="text-center text-3xl font-mono leading-6 mb-12 text-light-50">登陆</h3>
        <el-form-item prop="userName">
          <el-input v-model="loginFromData.userName" placeholder="用户名" suffix-icon="user"></el-input>
        </el-form-item>
        <el-form-item prop="passWord">
          <el-input v-model="loginFromData.passWord" placeholder="密码" suffix-icon="lock" show-password type="password"></el-input>
        </el-form-item>
        <el-button class="w-full" type="primary">登录</el-button>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';

defineOptions({
    name: "Login"
})
//   登录 
const validateUserName = (rule, value, callback) => {
  console.log("validateUserName",value,rule)
  if (value.length < 5) {
    return callback(new Error("输入正确的用户名"))
  }else {
    return callback()
  }
}
const validatePassword = (rule, value, callback) => {
  console.log("validatePassword",value, rule)
  if (value.length < 2) {
    return callback(new Error("输入正确的密码"))
  }else{
    callback()
  }
}
const loginFromData = reactive({
  userName: "",
  passWord: "",
})
const loginFromRules = reactive({
  userName:[{validator: validateUserName, trigger: "blur"}],
  passWord:[{validator: validatePassword, trigger: "blur"}],
})
const loginEntry = () => {}
</script>

<style>

</style>