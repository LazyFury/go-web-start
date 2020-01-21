<template>
  <div>
    <a-table
      :columns="columns"
      :rowKey="record => record.id"
      :dataSource="data"
      :pagination="pagination"
      :loading="loading"
      @change="handleTableChange"
    >
      <template slot="name" slot-scope="name">
        {{ name }}
      </template>
      <template slot="email" slot-scope="email">
        {{ email || "未绑定" }}
      </template>
      <template slot="status" slot-scope="status">
        <a-switch
          checkedChildren="正"
          unCheckedChildren="冻"
          :defaultChecked="Boolean(status)"
        />
      </template>

      <span slot="action" slot-scope="text, record">
        <a href="javascript:;">邀请 一 {{ record.name }}</a>
        <a-divider type="vertical" />
        <a href="javascript:;">删除</a>
        <a-divider type="vertical" />
        <a href="javascript:;" class="ant-dropdown-link">
          更多操作<a-icon type="down" />
        </a>
      </span>
    </a-table>
  </div>
</template>
<script>
import reqwest from "reqwest";
const columns = [
  {
    title: "姓名",
    dataIndex: "name",
    sorter: true,
    width: "20%",
    scopedSlots: { customRender: "name" }
  },
  {
    title: "IP",
    dataIndex: "ip",
    // filters: [{ text: 'Male', value: 'male' }, { text: 'Female', value: 'female' }],
    width: "14%"
  },
  {
    title: "最后登陆时间",
    dataIndex: "login_time",
    sorter: true,
    width: "20%"
  },
  {
    title: "Email",
    dataIndex: "email",
    scopedSlots: { customRender: "email" },
    width: "14%"
  },
  {
    title: "状态",
    dataIndex: "status",
    width: "5%",
    scopedSlots: { customRender: "status" }
  },
  {
    title: "Action",
    key: "action",
    scopedSlots: { customRender: "action" }
  }
];

export default {
  mounted() {
    this.fetch(this.pagination);
  },
  data() {
    return {
      data: [],
      pagination: {
        results: 10,
        current: 1
      },
      loading: false,
      columns
    };
  },
  methods: {
    handleTableChange(pagination, filters, sorter) {
      console.log(pagination);
      const pager = { ...this.pagination };
      pager.current = pagination.current;
      this.pagination = pager;
      this.fetch({
        results: pagination.results,
        current: pagination.current,
        sortField: sorter.field,
        sortOrder: sorter.order,
        ...filters
      });
    },
    fetch(params = {}) {
      console.log("params:", params);
      this.loading = true;
      reqwest({
        url: "http://127.0.0.1:8080/admin/user/list",
        method: "get",
        data: {
          page: params.current,
          limit: params.results
        },
        type: "json"
      })
        .then(res => {
          console.log(res);
          if (res.code == 1) {
            const pagination = { ...this.pagination };
            pagination.total = res.data.count;

            this.data = res.data.list;
            this.pagination = pagination;
          } else {
            this.$message.error(res.msg);
          }

          // Read total count from server
          // pagination.total = data.totalCount;
          this.loading = false;
        })
        .catch(err => {
          this.loading = false;
        });
    }
  }
};
</script>

<style scoped></style>
