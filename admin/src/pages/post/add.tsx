import useRequest from '@/hooks/useRequest';
import { postCates, posts } from '@/server/api/posts';
import { Button, Form, Input, PageHeader, Select } from 'antd';
import TextArea from 'antd/lib/input/TextArea';
import React, { useState } from 'react';
// import 'react-quill/dist/quill.snow.css';
import { history, useLocation } from 'umi';
import './add.less';

const layout = {
  labelCol: { span: 2 },
  wrapperCol: { span: 12 },
};
const { Option } = Select;
export default () => {
  const [content, setContent] = useState('');

  let param: any = useLocation();
  let { id } = param.query;
  let [isEdit] = useState(Boolean(id));
  let { data: cate, load: loadCate } = useRequest(postCates.list, true);
  const [form] = Form.useForm();

  const onFinish = (values: any) => {
    (() => {
      if (isEdit) {
        return posts.update(id, { ...values, content });
      }
      return posts.add({ ...values, content, tag: values.tag.split(',') });
    })().then(() => {
      history.push('/post');
    });
  };
  const onReset = () => {};

  return (
    <div>
      <PageHeader
        className="site-page-header fff"
        title={isEdit ? '修改文章' : '发布文章'}
        subTitle=""
      />
      <Form {...layout} form={form} onFinish={onFinish}>
        <Form.Item name="title" label="文章标题" rules={[{ required: true }]}>
          <Input />
        </Form.Item>

        <Form.Item
          labelCol={{ span: 2 }}
          wrapperCol={{ span: 4 }}
          name="cate_id"
          label="文章分类"
        >
          <Select allowClear placeholder="请选择文章分类">
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

        <Form.Item name="desc" label="文章简介" rules={[{ required: true }]}>
          <TextArea />
        </Form.Item>

        <Form.Item
          name="author"
          wrapperCol={{ span: 4 }}
          label="文章作者"
          rules={[{ required: true }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="email"
          wrapperCol={{ span: 4 }}
          label="作者邮箱"
          rules={[{ required: true }]}
        >
          <Input />
        </Form.Item>

        {/* <Row style={{ marginBottom: '20px' }}>
          <Col span={2} style={{ textAlign: 'right' }}>
            <text>文章内容:</text>
          </Col>
          <Col span={12} style={{ marginLeft: '10px' }}>
            <ReactQuill theme="snow" value={content} onChange={setContent} />
          </Col>
        </Row> */}

        <Form.Item name="tag" label="标签" rules={[{ required: true }]}>
          <Input />
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
