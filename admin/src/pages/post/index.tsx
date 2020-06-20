import { useDataList } from '@/hooks';
import { http } from '@/server/request';
import { PageHeader, Table } from 'antd';
import React from 'react';
const getList = (page: number) => {
  return http
    .get('http://127.0.0.1:8080/post', {
      params: { page },
      headers: { 'Content-Type': ' application/json' },
    })
    .then((res) => res.data);
};

export default () => {
  let { data, load } = useDataList(getList);
  return (
    <div>
      <PageHeader
        className="site-page-header fff"
        // onBack={() => null}
        title="文章管理"
        subTitle="This is a subtitle"
      />
      <Table
        columns={columns}
        dataSource={data.list}
        size={'large'}
        bordered={true}
        rowKey={'ID'}
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
  { title: 'Title', key: 'title', dataIndex: 'title' },
  {
    title: '简介',
    key: 'desc',
    dataIndex: 'desc',
    render(desc: any) {
      return <div>{desc || '暂无内容....'}</div>;
    },
  },
  { title: '创建时间', key: 'created_at', dataIndex: 'CreatedAt' },
  { title: '更新时间', key: 'update_at', dataIndex: 'UpdatedAt' },
  { title: '作者', key: 'author', dataIndex: 'author' },
  { title: 'tag', key: 'tag', dataIndex: 'tag' },
  { title: '操作', key: 'action' },
];
