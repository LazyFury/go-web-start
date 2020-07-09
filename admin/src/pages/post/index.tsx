import { useDataList } from '@/hooks';
import { posts } from '@/server/api';
import config from '@/utils/config';
import {
  DeleteOutlined,
  EditOutlined,
  PaperClipOutlined,
  SyncOutlined,
} from '@ant-design/icons';
import { Button, Modal, PageHeader, Table, Tooltip } from 'antd';
import React from 'react';
import { Link } from 'umi';

let { confirm } = Modal;
let resetTableData: () => Promise<void>; //在其他组件中使用重置列表

export default () => {
  let { data, load, loading } = useDataList(page => posts.list({ page }));
  resetTableData = load;
  return (
    <div>
      <PageHeader
        style={{ padding: '20rpx 0' }}
        // onBack={() => null}
        title="文章管理"
        subTitle={`共${data.count}篇文章，${data.pageCount}页，当前${data.pageNow}页`}
      />

      <div style={{ margin: '10px 0' }}>
        <Button onClick={() => load()}>
          <SyncOutlined></SyncOutlined>
          <span>刷新</span>
        </Button>
      </div>

      <Table
        columns={columns}
        dataSource={data.list}
        size={'large'}
        bordered={true}
        rowKey={'id'}
        loading={loading}
        pagination={{
          position: ['bottomLeft'],
          defaultCurrent: data.pageNow,
          total: data.count,
          showSizeChanger: false,
          onChange: load,
        }}
        rowSelection={{
          type: 'checkbox',
          onChange: (selectedRowKeys, selectedRows) => {
            console.log(
              `selectedRowKeys: ${selectedRowKeys}`,
              'selectedRows: ',
              selectedRows,
            );
          },
        }}
      ></Table>
    </div>
  );
};

// 表格列配置
const columns = [
  { title: 'ID', key: 'id', dataIndex: 'id' },
  { title: '文章标题', key: 'title', dataIndex: 'title', render: title },
  { title: '分类', key: 'cate', dataIndex: 'cate_id' },
  // {
  //   title: '简介',
  //   key: 'desc',
  //   dataIndex: 'desc',
  //   render(desc: any) {
  //     return <div>{desc || '暂无内容....'}</div>;
  //   },
  // },
  { title: '作者', key: 'author', dataIndex: 'author' },
  { title: '文章标签', key: 'tag', dataIndex: 'tag' },
  { title: '创建时间', key: 'created_at', dataIndex: 'created_at' },
  { title: '更新时间', key: 'updated_at', dataIndex: 'updated_at' },
  {
    title: '操作',
    key: 'action',
    dataIndex: 'id',
    render: (id: any) => {
      return (
        <>
          {edit(id)}
          {del(id)}
        </>
      );
    },
  },
];

function del(id: number) {
  return (
    <Tooltip title="删除文章">
      <a
        onClick={() => confirmDel(id)}
        style={{ marginLeft: '4px', color: '#d33', padding: '0 10px' }}
      >
        <span>删除</span>
      </a>
    </Tooltip>
  );
}

function edit(id: number) {
  return (
    <Tooltip title="编辑文章">
      <Link to={'/post/add?type=edit&id=' + id} style={{ marginLeft: '4px' }}>
        <EditOutlined></EditOutlined>
        <span>编辑</span>
      </Link>
    </Tooltip>
  );
}

function title(article: React.ReactNode) {
  return (
    <div>
      <span> {article}</span>
      <Tooltip title="预览文章">
        <a
          href={config.previewUrl}
          target="_blank"
          style={{ marginLeft: '4px' }}
        >
          <PaperClipOutlined></PaperClipOutlined>
        </a>
      </Tooltip>
    </div>
  );
}

function confirmDel(id: number) {
  return confirm({
    title: '确定要删除文章？无法撤回',
    icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
    okText: '确定删除',
    okType: 'danger',
    onOk() {
      posts.del(id).then(res => {
        resetTableData();
      });
    },
  });
}
