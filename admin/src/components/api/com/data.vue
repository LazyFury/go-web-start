<template>
  <div>
    <div>
      <div class="row config-line" v-for="(item, index) in tmpList" :key="index">
        <div width class="title hide-text-1">
          <span for>{{item.name}}</span>
        </div>
        <div style="flex:1;margin:0 10px;">
          <template v-if="item.type=='number'">
            <a-input-number id="inputNumber" v-model="item.value" @change="onChange($event,index)" />
          </template>
          <template v-else>
            <a-input v-model="item.value"></a-input>
          </template>
        </div>
        <div v-if="$isDev">
          <a href="javascript:;" @click="editLine(index)">编辑</a>
          /
          <a href="javascript:;" @click="delLine(index)">删除</a>
        </div>
      </div>
      <div>
        <a-button @click="showDrawer" v-if="$isDev">
          添加参数
          <a-icon type="plus"></a-icon>
        </a-button>
      </div>
    </div>

    <add-param ref="addParam" @save="addParamSave"></add-param>
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
      visible: false,
      isAdd: true //1 add 0 edit
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
    showDrawer() {
      this.$refs.addParam.show();
    },
    onChange(e, i) {
      console.log(e, i);
      this.tmpList[i].value = e;
      this.$emit("update", this.tmpList);
    },
    addParamSave(values) {
      console.log(values);

      if (!this.isAdd) {
        this.tmpList.splice(this.ChooseIndex, 1);
      }

      let { key } = values;
      console.log(key);
      for (let i = 0; i < this.tmpList.length; i++) {
        const element = this.tmpList[i];
        if (key == element.key) {
          this.$message.info("不可使用重复的参数Key");
          return;
        }
      }

      this.tmpList.push(values);
      this.isAdd = true;
    },
    delLine(i) {
      this.tmpList.splice(i, 1);
    },
    editLine(i) {
      this.ChooseIndex = i;
      this.isAdd = false;
      this.$refs.addParam.show(this.tmpList[i]);
    }
  }
};
</script>

<style>
.row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}
.config-line {
  margin: 8px 0;
  background: #eee;
  padding: 4px 10px;
}
.config-line .title {
  flex: 0 0 80px;
}

.hide-text-1 {
  display: -webkit-box !important;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  overflow: hidden;
}
</style>