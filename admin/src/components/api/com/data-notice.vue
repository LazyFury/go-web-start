<template>
  <div>
    <a-table
      :columns="columns"
      :dataSource="data && data.body || []"
      bordered
      style="margin-bottom:10px"
    >
      <template slot="title" slot-scope="json">Body 配置：{{json}}</template>
      <!-- <template slot="footer" slot-scope="currentPageData">Footer{{currentPageData}}</template> -->
    </a-table>
    <a-table :columns="columns" :dataSource="data && data.headers || []" bordered>
      <template slot="title" slot-scope="json">Header 配置：{{json}}</template>
      <!-- <template slot="footer" slot-scope="currentPageData">Footer{{currentPageData}}</template> -->
    </a-table>
  </div>
</template>

<script>
const columns = [
  {
    title: "名称",
    dataIndex: "name",
    scopedSlots: { customRender: "name" }
  },
  {
    title: "参数",
    // className: "column-money",
    dataIndex: "key"
  },
  {
    title: "值",
    dataIndex: "value"
  },
  {
    title: "类型",
    dataIndex: "type"
  },
  {
    title: "提示",
    dataIndex: "tips"
  }
];

export default {
  props: {
    data: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      columns
    };
  }
};
</script>

<style>
.ant-table-pagination {
  display: none !important;
}
</style>