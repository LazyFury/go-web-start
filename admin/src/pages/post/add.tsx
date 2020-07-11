import useRequest from '@/hooks/useRequest';
import { postCates, posts } from '@/server/api';
import { Button, Form, Input, PageHeader, Select } from 'antd';
import React, { useState } from 'react';
import { history, useLocation } from 'umi';
import './add.less';
const layout = {
  labelCol: { span: 2 },
  wrapperCol: { span: 12 },
};

export default () => {
  let param: any = useLocation();
  let { id } = param.query;
  let [isEdit] = useState(Boolean(id));
  let { data: cate, load: loadCate } = useRequest(postCates.list, true);
  const [form] = Form.useForm();
  const { Option } = Select;

  const onFinish = (values: any) => {
    if (isEdit) {
      posts.update(id, values);
      return;
    }
    posts.add(values);
  };
  const onReset = () => {};

  return (
    <div>
      <PageHeader
        className="site-page-header fff"
        onBack={() => history.go(-1)}
        title={isEdit ? '修改文章' : '发布文章'}
        subTitle=""
      />
      <Form {...layout} form={form} onFinish={onFinish}>
        <Form.Item name="title" label="文章标题" rules={[{ required: true }]}>
          <Input />
        </Form.Item>

        <Form.Item labelCol={{span:2}} wrapperCol={{span:4}} name="cate_id" label="文章分类">
          <Select
            allowClear
            placeholder="请选择文章分类"
          >
            {cate && cate.length > 0
              ? cate.map(
                  (x: { id: React.ReactText; name: React.ReactNode }) => {
                    return (
                      <Option key={x.id} value={x.id}>
                        {x.name}
                      </Option>
                    );
                  },
                )
              : null}
          </Select>
        </Form.Item>
        <Form.Item className="submit" wrapperCol={{ offset: 2, span: 16 }}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
          <Button htmlType="button" onClick={onReset}>
            Reset
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
