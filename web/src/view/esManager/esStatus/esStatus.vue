<template>
    <div>
        <el-tabs type="card">
            <el-tab-pane label="节点信息">
                <!-- <menu-permission ref="menus" :row="activeRow"></menu-permission> -->
            </el-tab-pane>
            <el-tab-pane label="集群索引信息">
                <!-- <api-permission ref="apis" :row="activeRow"></api-permission> -->
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script>
export default {
    name: "esStatus"
}
</script>

<style>

</style>