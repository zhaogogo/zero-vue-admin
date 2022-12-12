<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App"
}
</script>

<style lang="scss">
@import "@/styles/main.scss";
@import '@/styles/base.scss';
#app {
  height: 100vh;
  overflow: hidden;
  font-weight: 400;
}
</style>