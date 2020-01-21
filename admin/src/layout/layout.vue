<template>
  <a-layout class="layout">
    <header-custom></header-custom>
    <a-layout>
      <slide/>
      <a-layout style="padding: 0 24px 24px">
        <!-- <a-breadcrumb style="margin: 16px 0">
          <a-breadcrumb-item>Home</a-breadcrumb-item>
          <a-breadcrumb-item>List</a-breadcrumb-item>
          <a-breadcrumb-item>App</a-breadcrumb-item>
        </a-breadcrumb> -->
        <a-layout-content
          :style="{ background: '#fff', padding: '24px', margin:'20px 0', minHeight: '280px' }"
        >
          <slot></slot>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>
<script>
import slide from './slide'
import headerCustom from './header'
  export default {
    components:{
      slide,headerCustom
    },
    data() {
      return {
        collapsed: false,
      };
    },
  };
</script>

<style>
  .layout .logo {
    width: 120px;
    height: 31px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px 28px 16px 0;
    float: left;
    color: #fff;
    font-size: 30px
  }

  .row{
    display: flex;
    flex-direction: row
  }
</style>
