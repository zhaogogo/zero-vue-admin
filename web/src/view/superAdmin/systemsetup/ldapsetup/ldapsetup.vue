<template>
  <div class="ldapsetup">
    <p>ldap配置</p>
    <el-form ref="ldapSetupForm" :label-position="labelPosition" :model="ldapSetup" title="ldap配置" label-width="80px">
        <el-form-item label="ldap地址" prop="addr">
            <el-input v-model="ldapSetup.addr"></el-input>
        </el-form-item>
        <el-form-item label="baseDN" prop="baseDN">
            <el-input v-model="ldapSetup.baseDN"></el-input>
        </el-form-item>
        <el-form-item label="bindDN" prop="bindDN">
            <el-input v-model="ldapSetup.bindDN"></el-input>
        </el-form-item>
        <el-form-item label="password" prop="password">
            <el-input v-model="ldapSetup.password"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="text">测试ldap连接</el-button>
            <el-button type="primary">立即创建</el-button>
            <el-button>取消</el-button>
        </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
    name: "LdapSetup",
    data() {
        return {
            labelPosition: "right",
            ldapSetup: {
                addr: "",
                baseDN: "",
                bindDN: "",
                password: "",
                userfield: ""
            }
        }
    }
}
</script>

<style lang="scss">
.ldapsetup{
    margin: 30px;
    p {
        text-align: center;
        line-height: 20px;
        font-size: 25px;
        padding-bottom: 30px;
        color: #606266
;
    }
}
</style>