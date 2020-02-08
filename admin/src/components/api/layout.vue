<template>
  <div>
    <a-layout id="components-layout-demo-top-side-2">
      <a-layout-header>
        <div class="logo">API 管理工具</div>
      </a-layout-header>

      <a-layout style="min-height:calc(100vh - 64px)">
        <a-layout-sider style="background: #fff">
          <slot name='apilist'>
              list
          </slot>
        </a-layout-sider>
        <a-col style="width:300px;background:#f9f9f9">
            <slot name='config'>
              <div class="config">config</div>
          </slot>
        </a-col>
        <a-layout>
          <a-layout-content>
            <slot name="result"></slot>
          </a-layout-content>
          <a-layout-footer>- API 管理列表</a-layout-footer>
        </a-layout>
      </a-layout>
    </a-layout>
  </div>
</template>

<style scoped>
.logo{
  color: #fff;
  font-size: 20px
}
</style>