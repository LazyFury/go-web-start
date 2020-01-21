<template>
  <div>
      <layout id="app" v-if="isLayout">
        <router-view/>
      </layout>
       <router-view v-else/>
  </div>
</template>

<script>
import layout from '@/layout/layout.vue'

// 无需默认layout的页面
let noLayoutList = ['/login/login']
export default {
  name: 'App',
  components: {
    layout
  },
  data(){
    return {
      isLayout:true
    }
  },
  created(){
    let that = this
    console.log(this.$router)
    that.isLayout = !noLayoutList.includes(this.$router.currentRoute.path)

    this.$router.beforeEach((to,form,next)=>{
      console.log(to)
      if(noLayoutList.includes(to.path)){
        that.isLayout = false
      }else{
        that.isLayout = true
      }
      // console.log(that)
      next()
    })
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  height: 100vh;
}
</style>
