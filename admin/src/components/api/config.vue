<template>
  <div class="config">
    <a-col class="line">
      <label for>请求地址:</label>
      <a-input placeholder="Basic usage" v-model="result.url" />
    </a-col>
    <a-row class="line">
      <label for>请求方式:</label>
      <method :list="['POST','GET']" :value="result.method || ''" @change="handleChange"></method>
    </a-row>
    <a-row class="line">
      <label for>参数</label>
      <edit-data v-model="result.body">{{item}}</edit-data>
    </a-row>
    <a-row class="line">
      <label for>headers</label>
      <edit-data v-model="result.headers">{{item}}</edit-data>
    </a-row>
  </div>
</template>

<script>
import { http } from "../../server/api";
import method from "../../components/api/com/method";
import editData from "./com/data";
export default {
  components: {
    method,
    editData
  },
  props: {
    item: {
      type: Object,
      default: () => {
        return {};
      },
      deep: true
    }
  },
  data() {
    return {
      result: {}
    };
  },
  computed: {},
  model: {
    prop: "item",
    event: "update"
  },
  watch: {
    // 数据变动
    result: function(newVal, oldVal) {
      console.log("我是子组件，现在的值为：", newVal);
      this.$emit("update", newVal);
    },
    // props 更新
    item: function(newVal, oldVal) {
      this.result = newVal;
    }
  },
  methods: {
    handleChange(value) {
      this.update({ method: value });
    },
    update(data = {}) {
      let result = { ...this.result, ...data };
      this.result = result;
      this.$emit("update", result);
    }
  }
};
</script>

<style>
.config {
  padding: 20px;
  min-width: 300px;
}
.line {
  padding-bottom: 10px;
}
.line label {
  display: inline-block;
  margin-bottom: 5px;
  font-weight: bold;
  background: #eee;
  padding: 0 8px;
  line-height: 24px;
}
</style>