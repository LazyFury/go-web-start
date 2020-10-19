import PageMain from '@/components/PageMain';
import { Uploader } from '@/components/Upload/Upload';
import useRequest from '@/hooks/useRequest';
import { adEvents, adGroups, ads } from '@/server/api/ad';
import { Button, Form, Input, Select } from 'antd';
import React from 'react';

const { Option } = Select;

const layout: { labelCol: { span: number }; wrapperCol: { span: number } } = {
  labelCol: { span: 4 },
  wrapperCol: { span: 12 },
};

const AddAd = (props: { id?: number; onsubmit: () => void }) => {
  let { data: events, load: loadEvent, loading: eventsLoading } = useRequest(
    adEvents.list,
    true,
  );

  let { data: groups, load: loadGroup, loading: groupLoading } = useRequest(
    adGroups.list,
    true,
  );

  const formFinish = (e: any) => {
    ads.add(e).then(res => {
      props.onsubmit();
    });
  };

  return (
    <PageMain title="添加广告" subTitle="banner,海报">
      <Form
        {...layout}
        onFinish={formFinish}
        initialValues={{ group_id: props.id }}
      >
        <Form.Item name="title" label="标题" rules={[{ required: true }]}>
          <Input placeholder="请输入标题" />
        </Form.Item>
        <Form.Item name="param" label="参数" rules={[{ required: true }]}>
          <Input placeholder="选择参数" />
        </Form.Item>
        <Form.Item
          name="event_id"
          label="选择事件"
          rules={[{ required: true }]}
        >
          <Select
            allowClear
            placeholder="请选择广告位点击事件"
            onFocus={loadEvent}
            loading={eventsLoading}
          >
            {events && events.length > 0
              ? events.map((x: any) => {
                  return (
                    <Option value={x.id} key={x.event}>
                      {x.desc}
                    </Option>
                  );
                })
              : null}
          </Select>
        </Form.Item>

        <Form.Item
          name="group_id"
          label="选择广告位"
          rules={[{ required: true }]}
        >
          <Select
            allowClear
            placeholder="请选择广告位"
            onFocus={() => loadGroup()}
            loading={groupLoading}
            defaultValue={props.id}
            disabled={Boolean(props.id)}
          >
            {groups && groups.length > 0
              ? groups.map((x: any) => {
                  return (
                    <Option value={x.id} key={x.id + '-group'}>
                      {x.name}
                    </Option>
                  );
                })
              : null}
          </Select>
        </Form.Item>
        <Form.Item name="img" label="上传图片" rules={[{ required: true }]}>
          <Uploader />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 4 }}>
          <Button htmlType="submit" type="primary">
            提交
          </Button>
        </Form.Item>
      </Form>
    </PageMain>
  );
};
export default AddAd;
