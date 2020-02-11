<template>
  <div>
    <a-button type="primary" html-type="submit" @click="showDrawer" style="margin:20px;width:80%">
      添加API
      <a-icon type="plus" />
    </a-button>
    <a-menu mode="inline" style="height:calc(100vh - 140px);overflow-y:auto;overflow-x:hidden;">
      <a-sub-menu v-for="(sub,index) in list" :key="index">
        <span slot="title">{{sub.name}}</span>
        <a-menu-item v-for="(item,i) in sub.list" @click="change(item)" :key="i">{{item.name}}</a-menu-item>
      </a-sub-menu>
    </a-menu>

    <template>
      <div>
        <a-drawer
          title="添加API"
          placement="left"
          width="500"
          :closable="true"
          @close="onClose"
          :visible="visible"
        >
          <!-- api分类 -->
          <a-button @click="visible1=true">添加分类</a-button>
          <a-drawer
            title="添加API分类"
            placement="left"
            width="500"
            :closable="true"
            @close="onClose1"
            :visible="visible1"
          >
            <add-cate></add-cate>
          </a-drawer>
          <a-card style="margin-top:10px">{{addConfig}}</a-card>
          <!-- 配置 -->
          <config v-model="addConfig"></config>
          <a-button type="primary" @click="save" style="margin-left:20px">save</a-button>
        </a-drawer>
      </div>
    </template>
  </div>
</template>


<script>
import config from "./config";
import AddCate from "./com/AddCate";
export default {
  components: {
    config,
    AddCate
  },
  data() {
    return {
      visible: false,
      visible1: false,
      addConfig: {},
      list: [
        {
          name: "微信相关API",
          key: 1,
          list: [
            {
              name: "微信jsapiConfig授权",
              url: "/wechat/jsApiConfig",
              headers: [
                {
                  name: "Content-Type",
                  key: "Content-Type",
                  value: "application/json",
                  type: "string"
                }
              ],
              body: [
                {
                  name: "网址",
                  key: "url",
                  value: "http://baidu.com",
                  type: "string"
                }
              ],
              method: "GET"
            },
            {
              name: "123",
              url: "/wechat/wecj12at",
              headers: [
                {
                  name: "Content-Type",
                  key: "Content-Type",
                  value: "application/json",
                  type: "string"
                }
              ],
              body: [
                {
                  name: "ID",
                  key: "id",
                  value: 1,
                  type: "number"
                }
              ],
              method: "GET"
            }
          ]
        }
      ]
    };
  },

  methods: {
    change(item) {
      this.$emit("change", item);
    },
    showDrawer() {
      this.visible = true;
    },
    onClose() {
      this.visible = false;
    },
    onClose1() {
      this.visible1 = false;
    },
    save() {
      console.log(this.addConfig);
      this.list[0].list.push(JSON.parse(JSON.stringify(this.addConfig)));
      console.log(this.list);
      this.onClose();
    }
  }
};
</script>

<style>
</style>