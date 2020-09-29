import useRequest from '@/hooks/useRequest';
import { postCates, postTags } from '@/server/api/posts';
import { Button, Form, Input, PageHeader, Select } from 'antd';
import React from 'react';
const { Option } = Select;
const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 12 },
};

const AddPostTag = (props: { callback: () => void }) => {
  let { data: cate } = useRequest(postCates.list, true);

  const onFinish = (val: any) => {
    console.log(val);
    postTags.add(val).then(() => {
      props.callback instanceof Function && props.callback();
    });
  };
  return (
    <div className="page-main">
      <PageHeader
        style={{ padding: '20rpx 0' }}
        // onBack={() => null}
        title="添加标签"
        subTitle={``}
      />
      <Form {...layout} onFinish={onFinish}>
        <Form.Item name="val" label="文章标题" rules={[{ required: true }]}>
          <Input placeholder="标签名称"></Input>
        </Form.Item>
        <Form.Item name="cate_id" label="选择分类" rules={[{ required: true }]}>
          <Select placeholder="请选择分类">
            {cate instanceof Array &&
              cate.map(c => {
                return (
                  <Option key={c.id} value={c.id}>
                    {c.name}
                  </Option>
                );
              })}
          </Select>
        </Form.Item>

        <Form.Item className="submit" wrapperCol={{ offset: 6, span: 16 }}>
          <Button type="primary" htmlType="submit">
            添加标签
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default AddPostTag;
