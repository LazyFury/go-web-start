<template>
  <a-layout-sider width="200" style="background: #fff">
    <a-menu
      mode="inline"
      :defaultSelectedKeys="['1']"
      :defaultOpenKeys="['1']"
      :style="{ height: '100%', borderRight: 0 }"
    >
      <a-menu-item key="1">
        <router-link to="/">
          <a-icon type="home" />后台首页
        </router-link>
      </a-menu-item>
      <a-menu-item key="2">
        <router-link to="/api/list">
          <a-icon type="bar-chart" />接口列表
        </router-link>
      </a-menu-item>
      <a-sub-menu key="sub1">
        <span slot="title">
          <a-icon type="user" />用户管理
        </span>
        <a-menu-item key="5">
          <router-link to="/user/user">用户列表</router-link>
        </a-menu-item>

        <a-menu-item key="3">
          <router-link to="/user/invite">邀请关系</router-link>
        </a-menu-item>
        <a-menu-item key="4">
          <router-link to="/user/recharge">充值记录</router-link>
        </a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="sub2">
        <span slot="title">
          <a-icon type="laptop" />商品管理
        </span>
        <a-menu-item key="2">商品列表</a-menu-item>
        <a-menu-item key="6">商品分类</a-menu-item>
        <a-menu-item key="7">商家管理</a-menu-item>
        <a-menu-item key="8">库存核销</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="sub3">
        <span slot="title">
          <a-icon type="notification" />优惠活动
        </span>
        <a-menu-item key="9">优惠券</a-menu-item>
        <a-menu-item key="10">秒杀管理</a-menu-item>
        <a-menu-item key="11">团购管理</a-menu-item>
        <a-menu-item key="12">推广商品</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="sub4">
        <span slot="title">
          <a-icon type="highlight" />文章管理
        </span>
        <a-menu-item key="9">文章列表</a-menu-item>
        <a-menu-item key="9">添加文章</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="sub5">
        <span slot="title">
          <a-icon type="setting" />系统配置
        </span>
        <a-menu-item key="9">全局配置</a-menu-item>
        <a-menu-item key="9">插件配置</a-menu-item>
      </a-sub-menu>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {};
</script>

<style>
</style>