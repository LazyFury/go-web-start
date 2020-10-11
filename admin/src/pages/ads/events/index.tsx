import List from '@/components/List';
import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { adEvents } from '@/server/api/ad';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Drawer, Modal, Space, Tooltip } from 'antd';
import React, { useState } from 'react';
import { AddADEvent } from './components';

let { confirm } = Modal;

export default function AdEvents() {
  let { data, load, loading } = useRequest(adEvents.list, true);

  let [visible, setVisible] = useState(false);

  const deleteEvent = (id: number) => {
    confirm({
      title: '确定要删除广告位？无法撤回',
      icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
      okText: '确定删除',
      okType: 'danger',
      onOk() {
        adEvents.del(id).then(() => {
          load();
        });
      },
    });
  };

  const columns = [
    { title: 'ID', key: 'id', dataIndex: 'id' },
    { title: 'Event', key: 'event', dataIndex: 'event' },
    { title: 'desc', key: 'desc', dataIndex: 'desc' },
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
    {
      title: '操作',
      key: 'action',
      dataIndex: 'id',
      render: (id: number, data: any) => {
        return (
          <Space style={{ width: 120 }}>
            <a
              onClick={() => {
                //   setSelectValues(
                //     Object.keys(data).map(x => {
                //       return { name: x, value: data[x] };
                //     }),
                //   );
                setVisible(true);
              }}
            >
              编辑
            </a>
            <span> / </span>
            {(() => {
              if (data.count > 0 || data.tag_count > 0) {
                return (
                  <Tooltip title="该广告位下有子内容，不可删除，请先清理广告">
                    <span>不可删除</span>
                  </Tooltip>
                );
              } else {
                return <a onClick={() => deleteEvent(data.id)}>删除</a>;
              }
            })()}
          </Space>
        );
      },
    },
  ];
  return (
    <PageMain
      title="事件管理"
      subTitle="广告位事件，需要与前端约定如何处理事件"
    >
      <Drawer width={500} visible={visible} onClose={() => setVisible(false)}>
        <AddADEvent></AddADEvent>
      </Drawer>

      <List
        onRefresh={load}
        loading={loading}
        leftActions={[
          <Button key="add" type="primary" onClick={() => setVisible(true)}>
            <PlusOutlined />
            添加事件
          </Button>,
        ]}
        table={{
          loading,
          columns,
          bordered: true,
          rowKey: 'id',
          dataSource: (data instanceof Array && data) || undefined,
        }}
      />
    </PageMain>
  );
}
