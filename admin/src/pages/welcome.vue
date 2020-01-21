<template>
  <div class="hello">
    <div>
       <strong>
      <code style="font-size:20px">{{ msg }}</code>
    </strong>
    </div>
    <hr>
    <a-row :gutter="16" class="card">
      <a-col :span="6"
        >
        <a-card title="新用户">
          <a-statistic title="Feedback" :value="1128" style="margin-right: 50px">
          <template v-slot:suffix>
            <a-icon type="like" />
          </template> </a-statistic
      >
        </a-card>
        </a-col>
      <a-col :span="6"
        >
        <a-card title="订单量">
          <a-statistic title="Unmerged" :value="93" valueClass="demo-class">
        <template v-slot:suffix>
          <span> / 100</span>
        </template>
      </a-statistic>
        </a-card>
        </a-col>
      <a-col :span="6"
        >
        <a-card title="交易额">
          <a-statistic
            title="Feedback"
            :value="11.28"
            :precision="2"
            suffix="%"
            :valueStyle="{color: '#3f8600'}"
            style="margin-right: 50px"
          >
            <template v-slot:prefix>
              <a-icon type="arrow-up" />
            </template>
          </a-statistic>
        </a-card>
        </a-col>
      <a-col :span="6"
        >
        <a-card title="月流水"><a-statistic title="Feedback" :value="'$1128.00'" style="margin-right: 50px">
          <template v-slot:suffix>
          </template> </a-statistic
      ></a-card>
        </a-col>
    </a-row>
  </div>
</template>

<script>
import { Button } from "ant-design-vue";
export default {
  name: "HelloWorld",
  components: {
    Button
  },
  data() {
    return {
      msg: "Welcome to the new admin panel！"
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
