<template>
  <div>
  <a-row class="method">
    <a-button 
        size='small' 
        v-for="(item, index) in list" :key="index" 
        @click="change(item)" 
        :class="{'ant-btn-primary':value==item}"
        >{{item}}</a-button>
  </a-row>
  </div>
</template>

<script>
export default {
  name: "method",
  props: {
    list: Array,
    value: String
  },
  watch: {
    value(newVal, oldVal) {
      console.log(newVal, oldVal);
    }
  },
  methods: {
    change(val) {
      this.$emit("change", val);
    }
  }
};
</script>

<style>
.method .active {
  color: red;
}
.method button{
    margin-right: 10px;
    padding: 0 10px
}
</style>