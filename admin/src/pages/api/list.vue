<template>
  <div>
    <layout>
      <template slot="apilist">
        <list @change="listCheck"></list>
      </template>
      <template slot="config">
        <config v-model="item"></config>
      </template>

      <div slot="result" style="padding:20px">
        <div style="background:#fff;padding:20px;margin-bottom:20px">
          
          <a-button type="primary" @click="Send(item)" html-type="submit">发送请求</a-button>
        </div>
        <h2>response：</h2>
        <pre class="result" v-html="result"></pre>
      </div>
      
    </layout>
  </div>
</template>

<script>
import highLight from "./highlight";
import layout from "../../components/api/layout";
import list from "../../components/api/list";
import config from "../../components/api/config";
import axios from "axios";
let http = axios.create({
  baseURL: "http://0.0.0.0:8080"
});

export default {
  components: {
    layout,
    config,
    list
  },
  data() {
    return {
      title: "hello api!",
      item: {},
      result: "{}"
    };
  },
  methods: {
    listCheck(item) {
      console.log(item);
      this.item = item;
    },
    configChange(item) {
      this.item = item;
      this.$forceUpdate();
    },
    Send(item) {
      let { body, headers, method, url } = item;

      let data = this.getData(body);
      headers = this.getData(headers);

      if(!(data || headers)){
        this.$message.info("参数错误")
        return
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
          return res.data
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
      if(!list || list.length==0){
        return false
      }
      let obj = {};
      console.log(list);
      list.forEach(x => {
        obj[x.key] = x.value;
      });

      return obj;
    }
  }
};
</script>

<style >
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
</style>