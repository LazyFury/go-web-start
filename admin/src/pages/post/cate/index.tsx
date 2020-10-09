import List from '@/components/List';
import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { postCates } from '@/server/api/posts';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Drawer, Modal, Tooltip } from 'antd';
import React, { useState } from 'react';
import { AddPostCate } from './add';

let { confirm } = Modal;
export default function() {
  let { data, load, loading } = useRequest(postCates.list, true);
  let [values, setValues] = useState([{}]);
  let [visible, setVisible] = useState(false);

  function delCate(id: number) {
    confirm({
      title: '确定要删除文章？无法撤回',
      icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
      okText: '确定删除',
      okType: 'danger',
      onOk() {
        postCates.del(id).then(() => {
          load();
          setValues([]);
        });
      },
    });
  }

  function editCate(data: any = {}) {
    let { name = '', desc = '', key = '', id = null } = data;
    setValues([
      { name: 'id', value: id },
      { name: 'name', value: name },
      { name: 'desc', value: desc },
      { name: 'key', value: key },
    ]);
    setVisible(true);
  }

  return (
    <PageMain title="文章分类" subTitle={`共 ${data.length || 0} 文章分类`}>
      <Drawer width={500} visible={visible} onClose={() => setVisible(false)}>
        <AddPostCate
          onsubmit={() => {
            setVisible(false);
            load();
          }}
          values={values}
        ></AddPostCate>
      </Drawer>

      <div className="action-bar" style={{ margin: '10px 0' }}></div>
      <List
        onRefresh={load}
        loading={loading}
        leftActions={[
          <Button type="primary" onClick={editCate}>
            <PlusOutlined />
            <span>添加分类</span>
          </Button>,
        ]}
        table={{
          columns: [
            ...columns,
            {
              title: '操作',
              key: 'action',
              dataIndex: 'id',
              render: (id: number, data: any) => {
                return (
                  <div style={{ width: 120 }}>
                    <a onClick={() => editCate(data)}>编辑</a>
                    <span> / </span>
                    {(() => {
                      if (data.count > 0 || data.tag_count > 0) {
                        return (
                          <Tooltip title="该分类下有子文章或子标签，不可删除，请先清理文章">
                            <span>不可删除</span>
                          </Tooltip>
                        );
                      } else {
                        return <a onClick={() => delCate(id)}>删除</a>;
                      }
                    })()}
                  </div>
                );
              },
            },
          ],
          dataSource: data instanceof Array ? data : [],
          size: 'large',
          bordered: true,
          loading,
          rowKey: 'id',
          pagination: {
            total: data.length,
          },
        }}
      ></List>
    </PageMain>
  );
}

const columns = [
  { title: 'ID', key: 'id', dataIndex: 'id' },
  {
    title: '分类名称',
    key: 'name',
    dataIndex: 'name',
  },
  {
    title: '关键词',
    key: 'key',
    dataIndex: 'key',
  },
  { title: '描述', key: 'desc', dataIndex: 'desc' },
  { title: '文章统计', key: 'count', dataIndex: 'count' },
  { title: '标签统计', key: 'tag_count', dataIndex: 'tag_count' },
  {
    title: '更新时间',
    key: 'updated_at',
    dataIndex: 'updated_at',
  },
  {
    title: '创建时间',
    key: 'created_at',
    dataIndex: 'created_at',
  },
];
