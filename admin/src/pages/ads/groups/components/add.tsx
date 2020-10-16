import PageMain from '@/components/PageMain';
import { adGroups } from '@/server/api/ad';
import { Form, Input, InputNumber } from 'antd';
import Button from 'antd/es/button/button';
import React from 'react';

const layout: { labelCol: { span: number }; wrapperCol: { span: number } } = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

// 添加广告位;
export const AddAdGroup = (props: {
  callback: () => void;
  values: { name: string; value: any }[];
}) => {
  const formFinish = (e: any) => {
    const finish = () => {
      props.callback instanceof Function && props.callback();
    };
    let { value: id }: { value: any } =
      props.values?.filter(x => x.name == 'id')[0] || {};

    if (id) {
      adGroups.update(id, e).then(finish);
    } else {
      adGroups.add(e).then(finish);
    }
  };
  return (
    <PageMain
      title={props.values.length > 0 ? '修改广告位' : '添加广告位'}
      subTitle=""
    >
      <Form
        {...layout}
        className="add-form"
        onFinish={formFinish}
        fields={props.values}
      >
        <Form.Item name="name" label="广告位名称" rules={[{ required: true }]}>
          <Input type="text" placeholder="请输入广告位名称..." />
        </Form.Item>
        <Form.Item name="desc" label="介绍" rules={[{ required: true }]}>
          <Input type="text" placeholder="请输入介绍内容..." />
        </Form.Item>

        <Form.Item name="max_count" label="限制最大数量(默认：1)">
          <InputNumber value={3} />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 6 }}>
          <Button htmlType="submit" type="primary">
            提交
          </Button>
        </Form.Item>
      </Form>
    </PageMain>
  );
};
