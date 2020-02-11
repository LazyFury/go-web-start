<template>
  <div>
    <div>
      <table>
        <tbody>
          <tr v-for="(item, index) in tmpList" :key="index">
            <td width="60">
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
            <a-button @click="showDrawer">
              添加参数
              <a-icon type="plus"></a-icon>
            </a-button>
          </tr>
        </tbody>
      </table>
    </div>

    <a-drawer
      title="添加API参数"
      placement="left"
      width="500"
      :closable="true"
      @close="onClose"
      :visible="visible"
    >
      <add-param @save="addParamSave"></add-param>
    </a-drawer>
  </div>
</template>

<script>
import AddParam from "./AddParam";
export default {
  components: {
    AddParam
  },
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
      tmpList: [],
      visible: false
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
    onClose() {
      this.visible = false;
    },
    showDrawer() {
      this.visible = true;
    },
    onChange(e, i) {
      console.log(e, i);
      this.tmpList[i].value = e;
      this.$emit("update", this.tmpList);
    },
    addParamSave(e) {
      this.visible = false;
      console.log(e);
      this.tmpList.push(e);
    }
  }
};
</script>

<style>
</style>