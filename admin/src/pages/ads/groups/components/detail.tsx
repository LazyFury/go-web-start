import MyImage from '@/components/Image';
import List from '@/components/List';
import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { adGroups, ads } from '@/server/api/ad';
import { PlusOutlined } from '@ant-design/icons';
import { Button, Col, Drawer, Row, Space } from 'antd';
import React, { useEffect, useState } from 'react';
import AddAd from '../../add';

export const Detail = (props: { id: number }) => {
  let { data, load, loading } = useRequest(() => adGroups.detail(props.id));
  let [visible, setVisible] = useState(false);

  const del = (id: number) => {
    ads.del(id).then(() => {
      load();
    });
  };

  useEffect(() => {
    load();
  }, [props.id]);

  const columns = [
    { title: 'ID', key: 'id', dataIndex: 'id' },
    {
      title: '预览',
      key: 'id',
      width: 120,

      dataIndex: 'image',
      render: (url: string) => (
        <MyImage width={120} height={60} src={url}></MyImage>
      ),
    },
    { title: '标题', key: 'title', dataIndex: 'title' },
    { title: '参数', key: 'param', dataIndex: 'param' },
    { title: '事件', key: 'event', dataIndex: 'event' },
    {
      title: '操作',
      key: 'id',
      dataIndex: 'id',
      render: (_: any, record: any) => {
        return (
          <Space>
            <a>编辑</a>/<a onClick={() => del(record.id)}>删除</a>
          </Space>
        );
      },
    },
  ];
  return (
    <PageMain
      title={`广告位【${data.name}】`}
      subTitle={
        <>
          <Col>简介:{data.desc}</Col>

          <Row style={{ width: '80%' }}>
            <Space>
              <Col>限制数量:{data.max_count}</Col>
              <span>/</span>
              <Col>当前数量:{data.count}</Col>
            </Space>
          </Row>
        </>
      }
      loading={loading}
    >
      <Drawer width={1000} visible={visible} onClose={() => setVisible(false)}>
        <AddAd
          id={props.id}
          onsubmit={() => {
            setVisible(false);
            load();
          }}
        ></AddAd>
      </Drawer>

      <List
        loading={loading}
        onRefresh={load}
        leftActions={[
          <Button
            type="primary"
            key="add-ad-btn"
            onClick={() => setVisible(true)}
            disabled={data.count >= data.max_count}
          >
            <PlusOutlined />
            {data.count >= data.max_count
              ? '无法继续添加，限制最多' + data.max_count + '个'
              : '添加广告'}
          </Button>,
        ]}
        table={{
          dataSource: (data.list instanceof Array && data.list) || [],
          columns,
          bordered: true,
          rowKey: 'id',
        }}
      ></List>
    </PageMain>
  );
};
