<template>
  <layout id="app">
    <router-view/>
  </layout>
</template>

<script>
import layout from '@/layout/layout.vue'
export default {
  name: 'App',
  components: {
    layout
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  height: 100vh;
}
</style>
