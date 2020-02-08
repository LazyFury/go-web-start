<template>
  <div>
    <div >
      <table>
        <tbody>
          <tr v-for="(item, index) in tmpList" :key="index">
            <td width='60'>
              <span for>{{item.name}}</span>
            </td>
            <td>
              <template v-if="item.type=='number'">
                <a-input-number
                  id="inputNumber"
                  v-model="item.value"
                  @change="onChange($event,index)"
                />
              </template>
              <template v-else>
                <a-input v-model="item.value"></a-input>
              </template>
            </td>
          </tr>
          <tr>
            <a-button>添加参数
              <a-icon type='plus'></a-icon>
            </a-button>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "edit-data",
  props: {
    list: {
      type: Array,
      deep: true
    }
  },
  model: {
    prop: "list",
    event: "update"
  },
  data() {
    return {
      tmpList: []
    };
  },
  watch: {
    tmpList(val) {
      console.log(val);
      this.$emit("update", val);
    },
    list(val) {
      this.tmpList = val;
    }
  },
  methods: {
    onChange(e, i) {
      console.log(e, i);
      this.tmpList[i].value = e;
      this.$emit("update", this.tmpList);
    }
  }
};
</script>

<style>
</style>