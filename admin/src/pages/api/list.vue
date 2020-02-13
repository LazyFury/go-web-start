<template>
  <div>
    <layout>
      <template slot="apilist">
        <div style>
          <list ref="list" @change="listCheck"></list>
        </div>
      </template>
      <template slot="config">
        <div style="padding:20px;width:400px">
          <a-card style="margin-bottom:10px">
            <label for class="label">修改分类名称 （分类ID:{{item.cate.ID}}）</label>
            <a-input
              @blur="SaveCate"
              :disabled="!$isDev"
              placeholder="分类名称"
              v-model="item.cate.name"
            ></a-input>
            <!-- <a-button type="danger" style="margin-top:10px" @click="DelCate(item.cate.ID)">删除分类</a-button> -->
          </a-card>
          <a-card>
            <config v-model="item.data"></config>
            <a-button v-if="$isDev" type="primary" @click="UpdateConfig">更新配置</a-button>
            <a-button
              v-if="$isDev"
              type="danger"
              style="margin-left:10px"
              @click="delAPI(item.ID)"
            >删除API</a-button>
          </a-card>
        </div>
      </template>

      <div slot="result" style="padding:20px">
        <a-card class="card">
          <div style="background:#eee">
            <a-button type="primary" @click="Send(item)" html-type="submit">发送请求</a-button>
            <span>{{baseURL}}{{item.url||""}}</span>
          </div>
        </a-card>

        <h2>Response:</h2>
        <a-card>
          <pre class="result" style="max-height:400px;overflow-y:auto;" v-html="result"></pre>
        </a-card>

        <h2>Config:</h2>
        <a-card class="card">
          <config-notice :data="item && item.data"></config-notice>
        </a-card>
      </div>
    </layout>

    <a-modal
      title="Title"
      :visible="ModelVisible"
      @ok="ModelHandleOk"
      :confirmLoading="ModelConfirmLoading"
      @cancel="ModelVisible = false"
    >
      <p>确认要修改配置吗?</p>
    </a-modal>
    <a-modal
      title="Title"
      :visible="delAPIVisible"
      @ok="delAPIHandleOk"
      @cancel="delAPIVisible = false"
    >
      <p>确认要删除API吗?</p>
    </a-modal>
    <a-modal
      title="Title"
      :visible="delCateVisible"
      @ok="delCateHandleOk"
      @cancel="delCateVisible = false"
    >
      <p>确认要修改配置吗?</p>
    </a-modal>
  </div>
</template>

<script>
import highLight from "./highlight";
import layout from "../../components/api/layout";
import list from "../../components/api/list";
import config from "../../components/api/config";
import ConfigNotice from "../../components/api/com/data-notice";
import axios from "axios";
import globalConfig from "../../config";
let { baseURL } = globalConfig;
let http = axios.create({
  baseURL
});
import Vue from "vue";
import "./clipBoard";
import { GetParam } from "../../util/util";
export default {
  components: {
    layout,
    config,
    list,
    ConfigNotice
  },
  data() {
    return {
      baseURL,
      title: "hello api!",
      ModelVisible: false,
      ModelConfirmLoading: false,
      delCateVisible: false,
      delAPIVisible: false,
      item: {
        cate: {}
      },
      result: "{}"
    };
  },
  mounted() {
    let { isDev = "" } = this.$util.GetParam();
    isDev = Boolean(isDev);
    console.log(isDev);
    Vue.prototype.$isDev = isDev;
  },
  methods: {
    clipBoardText() {
      console.log("test");
    },
    listCheck(item) {
      console.log(item);
      this.item = item;
    },
    configChange(item) {
      this.item = item;
      this.$forceUpdate();
    },
    Send(item) {
      let { body, headers, method, url } = item.data;

      let data = this.getData(body);
      headers = this.getData(headers);

      if (!(data || headers || url)) {
        this.$message.info("参数错误");
        return;
      }

      let option = { data, headers, method, params: data, url };
      if (method == "POST") {
        delete option.params;
      }
      http
        .request(option)
        .then(res => {
          // 正常请求 200
          return res;
        })
        .catch(err => {
          // 异常http码
          return err.response;
        })
        .then(res => {
          return res.data;
        })
        .then(res => {
          console.log(res);
          if (res.code == 1) {
            this.$message.success(res.msg);
          } else {
            this.$message.error(res.msg);
          }

          this.result = highLight(JSON.stringify(res, undefined, 3));
        });
    },
    getData(list) {
      if (!list || list.length == 0) {
        return false;
      }
      let obj = {};
      console.log(list);
      list.forEach(x => {
        obj[x.key] = x.value;
      });

      return obj;
    },
    SaveCate(e) {
      setTimeout(() => {
        let { ID: id, name } = this.item.cate;
        if (!id) {
          return;
        }
        this.api.api.cate.save({ id, name }).then(res => {
          if (res.code == 1) {
            this.$refs.list.init();
          }
        });
      }, 300);
    },
    UpdateConfig() {
      if (!this.item.data) return;
      this.ModelVisible = true;
    },
    ModelHandleOk() {
      this.ModelConfirmLoading = true;
      let temp = JSON.parse(JSON.stringify(this.item));
      let { ID: id } = temp;

      let { name } = temp.data;
      delete temp.cate;
      let desc = JSON.stringify(temp.data);
      console.log(desc);

      this.api.api.api
        .save({ id, name, desc })
        .then(res => {
          if (res.code == 1) {
            this.$refs.list.init();
          }
        })
        .finally(res => {
          setTimeout(() => {
            this.ModelVisible = false;
            this.ModelConfirmLoading = false;
          }, 300);
        });
    },
    delAPI(id) {
      if (!this.item.data) return;
      this.delAPIVisible = true;
      this.waitDelAPIID = id;
    },
    delAPIHandleOk() {
      let id = this.waitDelAPIID;
      this.api.api.api
        .del({ id })
        .then(res => {
          if (res.code == 1) {
            this.$refs.list.init();
            this.item = { cate: {} };
          }
        })
        .finally(res => {
          this.delAPIVisible = false;
        });
    },
    DelCate(id) {
      if (!this.item.data) return;
      this.delCateVisible = true;
      this.waitDelCateID = id;
    },
    delCateHandleOk() {
      let id = this.waitDelCateID;
      this.api.api.cate
        .del({ id })
        .then(res => {
          if (res.code == 1) {
            this.$refs.list.init();
          }
        })
        .finally(res => {
          this.delCateVisible = false;
        });
    }
  }
};
</script>

<style >
.card {
  margin-bottom: 10px;
}
pre {
  outline: 1px solid #ccc;
  padding: 5px;
  margin: 5px;
  max-height: 80vh;
  overflow-y: auto;
}
.string {
  color: green;
}
.number {
  color: darkorange;
}
.boolean {
  color: blue;
}
.null {
  color: magenta;
}
.key {
  color: red;
}

.label {
  display: inline-block;
  margin-bottom: 5px;
  font-weight: bold;
  background: #eee;
  padding: 0 8px;
  line-height: 24px;
}
</style>