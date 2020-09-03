import { useDataList } from '@/hooks/useDataList';
import { postCates } from '@/server/api/posts';
import { Button, Drawer, Form, Input, PageHeader } from 'antd';
import React, { useState } from 'react';

export default function() {
  let { data, load, loading } = useDataList(postCates.list);

  let [visible, setVisible] = useState(false);

  return (
    <div>
      <PageHeader
        className="site-page-header fff"
        onBack={() => history.go(-1)}
        title="文章分类"
        subTitle=""
      />
      <Button onClick={() => setVisible(true)}>show</Button>

      <Drawer visible={visible} onClose={() => setVisible(false)}>
        <AddPostCate></AddPostCate>
      </Drawer>

      <div>{JSON.stringify(data)}</div>
    </div>
  );
}

function AddPostCate() {
  return (
    <Form>
      <Form.Item name="name" label="name" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item>
        <Button htmlType="submit">submit</Button>
      </Form.Item>
    </Form>
  );
}
