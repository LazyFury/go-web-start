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
      <template slot="status" slot-scope="status, record, index">
        <a-popconfirm
          :title="'是否要' + (Boolean(status) ? '冻结' : '解冻') + '该用户?'"
          @confirm="frozen(!Boolean(status), record, index)"
          okText="确定"
          cancelText="取消"
        >
          <a-switch
            :ref="'switch-' + record.id"
            checkedChildren="正"
            unCheckedChildren="冻"
            :checked="Boolean(status)"
          />
        </a-popconfirm>
      </template>

      <span slot="action" slot-scope="text, record">
        <router-link :to="'./edit?id=' + record.id">编辑</router-link>

        <a-divider type="vertical" />

        <a-dropdown>
          <a class="ant-dropdown-link" href="#">
            更多操作<a-icon type="down" />
          </a>
          <a-menu slot="overlay">
            <a-menu-item>
              <a href="javascript:;">用户充值</a>
            </a-menu-item>
            
            <a-menu-item>
              <a href="javascript:;">订单查询</a>
            </a-menu-item>

            <a-menu-item>
              <a href="javascript:;">推广查询</a>
            </a-menu-item>
          </a-menu>
        </a-dropdown>
        <a-divider type="vertical" />
        <a-popconfirm
          title="是否立即删除该用户?一旦删除无法恢复！"
          @confirm="del(record)"
          okText="确定删除"
          okType="danger"
          cancelText="取消"
          ><a href="javascript:;" style="color:gray">删除</a></a-popconfirm
        >
      </span>
    </a-table>
  </div>
</template>
<script>
import reqwest from "reqwest";
const columns = [
  {
    title: "用户id",
    dataIndex: "id",
  },
  {
    title: "姓名",
    dataIndex: "name",
    sorter: true,
    width: "20%",
    scopedSlots: { customRender: "name" }
  },
  {
    title: "安全邮箱",
    dataIndex: "email",
    scopedSlots: { customRender: "email" },
    width: "14%"
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
    title: "是否正常",
    dataIndex: "status",
    width: "8%",
    scopedSlots: { customRender: "status" }
  },
  {
    title: "操作",
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
  created() {
    // this.api.adminHome()
  },
  methods: {
    del({ id }) {
      this.api.user.del({ id }).then(res => {
        this.fetch(this.pagination);
      });
    },
    frozen(checked, item, index) {
      let status = Number(checked);
      let { id } = item;
      // console.log()
      this.api.user
        .frozen({ id, status })
        .then(res => {
          this.data[index].status = status;
        })
        .catch(err => {});
    },
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
      this.api.user
        .list({
          page: params.current,
          limit: params.results
        })
        .then(res => {
          console.log(res);
          if (res.code == 1) {
            const pagination = { ...this.pagination };
            pagination.total = res.data.count;

            this.data = res.data.list;
            this.pagination = pagination;
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
