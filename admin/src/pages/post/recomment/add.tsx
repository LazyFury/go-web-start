import { postRec } from '@/server/api/posts';
import { Button, Form, Input, PageHeader } from 'antd';
import React from 'react';
import './add.less';

export function AddPostCate(props: {
  onsubmit?: () => void;
  onerror?: () => void;
  values?: any[] | undefined;
}) {
  const onFinish = (values: any) => {
    let { value: id }: { value: number } = props.values?.filter(
      x => x.name == 'id',
    )[0];
    const finish = () => {
      props.onsubmit && props.onsubmit();
    };
    if (id) {
      return postRec.update(id, values).then(finish);
    }
    postRec.add(values).then(finish);
  };

  return (
    <div>
      <PageHeader title="添加推荐位"></PageHeader>
      <Form
        labelCol={{ span: 5 }}
        wrapperCol={{ span: 12 }}
        className="add-form"
        onFinish={onFinish}
        fields={props.values}
      >
        <Form.Item name="name" label="name" rules={[{ required: true }]}>
          <Input />
        </Form.Item>

        <Form.Item name="key" label="keywords" rules={[]}>
          <Input />
        </Form.Item>
        <Form.Item name="desc" label="desc" rules={[]}>
          <Input />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 5, span: 6 }}>
          <Button htmlType="submit" type="primary">
            submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}
