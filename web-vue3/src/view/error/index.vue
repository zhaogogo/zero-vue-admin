<template>
    <el-result 
        class="bg-light-100 h-screen" 
        title="OOPS!" 
        sub-title="页面无权限请联系管理员"
    >
        <template #icon>
            <el-image
                src="@/assets/404_image/404.png"
            ></el-image>
        </template>
        <template #extra>
            <el-button type="permary">Back</el-button>
        </template>
    </el-result>
</template>

<script setup>
defineOptions({
    name: "ErrorPage"
})
</script>

<style scoped>

</style>