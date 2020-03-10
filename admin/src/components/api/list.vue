<template>
  <div style>
    <a-button
      type="primary"
      v-if="$store.state.isDev"
      html-type="submit"
      @click="showDrawer"
      style="margin:20px;width:80%"
    >
      添加API
      <a-icon type="plus" />
    </a-button>
    <a-menu mode="inline" style="height:calc(100vh - 140px);overflow-y:auto;overflow-x:hidden;">
      <a-sub-menu v-for="(sub,index) in list" :key="index" @titleClick="titleClick($event,index)">
        <span slot="title">{{sub.name}}</span>
        <a-menu-item
          v-for="(item,i) in sub.list"
          @click="change(item,sub)"
          :key="index+'_'+i"
        >{{item.name||'没有'}}</a-menu-item>
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
          <a-button @click="visible1=true" type="primary">添加分类</a-button>
          <label for>选择分类:</label>
          <a-select style="width: 120px" @change="handleAPICateChange">
            <a-select-option v-for="(cate,index) in cateList" :key="index">{{cate.name}}</a-select-option>
          </a-select>

          <a-drawer
            title="添加API分类"
            placement="left"
            width="500"
            :closable="true"
            @close="onClose1"
            :visible="visible1"
          >
            <add-cate @save="cateSave"></add-cate>
          </a-drawer>
          <a-card style="margin-top:10px;margin-bottom:10px">
            <div>{{currentCate}}</div>
            <div>{{addConfig}}</div>
          </a-card>
          <!-- 配置 -->
          <config v-model="addConfig"></config>
          <a-button type="primary" @click="save" style="margin-left:0px;padding:0 20px">save</a-button>
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
      cateList: [{ name: "请选择" }],
      currentCate: null,
      list: []
    };
  },
  watch: {
    visible(val) {
      if (val) {
        this.GetCate();
      }
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    titleClick(e, index) {
      console.log(e);
      let { ID: cid } = this.list[index];
      // this.api.api.cate
      //   .api({ cid })
      //   .then(res => {
      //     this.list[index]["list"] = res.data;
      //   })
      //   .catch(err => {
      //     this.list[index]["list"] = [{ name: "加载失败" }];
      //   });
    },
    init() {
      Promise.all([
        this.api.api.api.all().then(res => {
          console.log(res.data);
          this.list = res.data.map(x => {
            // x["list"] = [];
            return x;
          });
        })
      ]);
    },
    GetCate() {
      return this.api.api.cate.GetAll().then(res => {
        this.cateList = res.data;
        return res.data;
      });
    },
    loadAPI() {
      console.log("click");
    },
    handleAPICateChange(i) {
      console.log(this.cateList[i]);
      this.currentCate = this.cateList[i];
    },
    change(item, cate) {
      let data = {};
      try {
        data = JSON.parse(item.data);
      } catch (err) {
        console.error("API配置内容解析错误 err:" + err);
      }
      let { ID, name } = cate;
      this.$emit("change", {
        ...item,
        data: { body: [], header: [], ...data },
        cate: { ID, name }
      });
    },
    showDrawer() {
      this.visible = true;
    },
    onClose() {
      this.visible = false;
      this.init();
    },
    onClose1() {
      this.visible1 = false;
    },
    save() {
      if (!this.currentCate) {
        this.$message.error("请选择分类先");
        return;
      }
      console.log(this.addConfig);
      this.api.api.api
        .add({
          name: this.addConfig.name,
          data: JSON.stringify(this.addConfig),
          cid: this.currentCate.ID
        })
        .then(res => {
          this.onClose();
        });
    },
    cateSave() {
      this.GetCate();
      this.onClose1();
    }
  }
};
</script>

<style>
</style>