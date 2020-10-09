import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { postCates, posts, postTags } from '@/server/api/posts';
import { emptyPromise } from '@/utils/utils';
import { Button, Col, Drawer, Form, Input, Row, Select, Spin } from 'antd';
import React, { useState } from 'react';
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';
import { history, useLocation } from 'umi';
import './add.less';
import AddPostTag from './tag/add';

const layout: { labelCol: { span: number }; wrapperCol: { span: number } } = {
  labelCol: { span: 4 },
  wrapperCol: { span: 12 },
};
const { Option } = Select;
export default () => {
  const [content, setContent] = useState('');

  let param: any = useLocation();
  let { id } = param.query;
  let [isEdit] = useState(Boolean(id));

  let { data, load, loading } = useRequest(() => {
    if (!id) {
      return emptyPromise();
    }
    return posts.detail(id).then(res => {
      if (res) {
        let { content } = res.data;
        setContent(content);
      }
      return res;
    });
  }, true);
  // 添加标签抽屉
  let [showAddTags, setShowAddTags] = useState(false);
  // 分类
  let { data: cate, load: loadCate } = useRequest(postCates.list, true);
  // tag
  let { data: tags, load: loadTags } = useRequest(postTags.list, true);
  const [form] = Form.useForm();
  const init = () => {
    loadCate();
    loadTags();
  };
  const onFinish = (values: any) => {
    console.log(values);
    (() => {
      if (isEdit) {
        return posts.update(id, { ...values, content });
      }
      return posts.add({ ...values, content });
    })().then(() => {
      history.push('/post');
    });
  };
  const onReset = () => {};

  return (
    <PageMain title={isEdit ? '修改文章' : '发布文章'} subTitle="">
      <Spin spinning={loading}>
        <div className="page-main">
          <Drawer
            width={500}
            visible={showAddTags}
            onClose={() => setShowAddTags(false)}
          >
            <AddPostTag
              callback={() => {
                setShowAddTags(false);
                init();
              }}
            ></AddPostTag>
          </Drawer>

          <Form
            {...layout}
            form={form}
            onFinish={onFinish}
            fields={Object.keys(data).map(key => {
              return {
                name: key,
                value: data[key],
              };
            })}
          >
            <Form.Item
              name="title"
              label="文章标题"
              rules={[{ required: true }]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              name="cate_id"
              label="文章分类"
              rules={[{ required: true }]}
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

            <Form.Item label="文章简介">
              <ReactQuill
                theme="snow"
                style={{ height: '300px' }}
                value={content || ''}
                onChange={setContent}
              />
            </Form.Item>

            <Form.Item
              name="author"
              label="文章作者"
              rules={[{ required: false }]}
            >
              <Input />
            </Form.Item>

            <Form.Item name="email" label="作者邮箱">
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
              <Select
                mode="multiple"
                style={{ width: '100%', marginRight: '10px' }}
                placeholder="选择文章标签..."
                optionLabelProp="label"
              >
                {tags instanceof Array &&
                  tags.map(tag => {
                    return (
                      <Option key={tag.id} value={tag.val}>
                        {tag.val || '~'}
                      </Option>
                    );
                  })}
              </Select>
            </Form.Item>

            <Row style={{ marginBottom: '20px', marginTop: '-10px' }}>
              <Col span={layout.wrapperCol.span} offset={layout.labelCol.span}>
                <Button type="dashed" onClick={() => setShowAddTags(true)}>
                  添加标签
                </Button>
              </Col>
            </Row>

            <Form.Item
              className="submit"
              wrapperCol={{
                offset: layout.labelCol.span,
                span: layout.wrapperCol.span,
              }}
            >
              <Button type="primary" htmlType="submit">
                Submit
              </Button>
              <Button htmlType="button" onClick={onReset}>
                Reset
              </Button>
            </Form.Item>
          </Form>
        </div>
      </Spin>
    </PageMain>
  );
};
