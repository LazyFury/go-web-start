<template>
  <div>
    <a-layout id="components-layout-demo-top-side-2">
      <a-layout-header>
        <a-row>
          <a-col :span="12">
            <div class="logo">API 管理工具</div>
          </a-col>
          <a-col :span="12" class="setting">
            <div class="menu">
              <a href="javascript:;">
                <a-icon type="setting" :style="{color:'#fff','font-size':'20px'}" />
              </a>
            </div>
          </a-col>
        </a-row>
      </a-layout-header>

      <a-layout style="height:calc(100vh - 64px);">
        <a-layout-sider class="sider" style>
          <slot name="apilist">list</slot>
        </a-layout-sider>
        <a-col style="width:auto;background:#f9f9f9;" class="hide-config">
          <slot name="config">
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
.logo {
  color: #fff;
  font-size: 20px;
}
.sider {
  background: #fff;
  flex: 0 0 240px !important;
  max-width: 240px !important;
  z-index: 99;
}
.hide-config {
  /* transform: translateX(-100%); */
}
.setting {
  justify-content: flex-end;
  align-items: center;
  display: flex;
}
</style>