import { adEvents } from '@/server/api/ad';
import { Button, Form, Input, PageHeader } from 'antd';
import React from 'react';

const layout: { labelCol: { span: number }; wrapperCol: { span: number } } = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

export const AddADEvent = () => {
  const formFinish = (e: any) => {
    adEvents.add(e);
  };
  return (
    <div className="page-main">
      <PageHeader title="添加事件"></PageHeader>

      <Form className="add-form" {...layout} onFinish={formFinish}>
        <Form.Item name="event" label="Event" rules={[{ required: true }]}>
          <Input placeholder="仅支持英文字符串" />
        </Form.Item>
        <Form.Item name="desc" label="事件介绍" rules={[{ required: true }]}>
          <Input placeholder="请输入事件介绍..." />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 8, span: 6 }}>
          <Button htmlType="submit" type="primary">
            提交
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
