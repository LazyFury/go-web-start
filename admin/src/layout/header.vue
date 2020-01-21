<template>
  <a-layout-header class="header">
    <a-row :style="{ 'justify-content': 'space-between' }">
      <a-col :span="12">
        <a-row>
          <a-col>
            <router-link to="/" class="logo"> </router-link>
          </a-col>
          <a-col :span="4">
            <a-menu
              theme="dark"
              mode="horizontal"
              :defaultSelectedKeys="['1']"
              :style="{ lineHeight: '64px' }"
              :selectable="false"
            >
              <!-- <a-menu-item key="2">
                <router-link to="/">
                  <a-icon type="home"></a-icon>
                  首页</router-link
                ></a-menu-item> -->
            </a-menu>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="12">
        <a-menu
          theme="dark"
          mode="horizontal"
          :defaultSelectedKeys="['1']"
          :style="{ lineHeight: '64px', 'text-align': 'end' }"
          :selectable="false"
        >
          <!-- <a-menu-item key="1">首页</a-menu-item> -->
          <a-menu-item key="2">
            消息通知（1）<a-badge dot>
      <a-icon type="notification" />
    </a-badge>
             </a-menu-item>
          <a-menu-item key="3">管理员设置<a-icon type="down" /></a-menu-item>
          <a-avatar icon="user" />
        </a-menu>
      </a-col>
    </a-row>
  </a-layout-header>
</template>

<script>
export default {};
</script>

<style></style>
