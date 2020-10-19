import { register } from '@/server/api/users';
import { LockOutlined, MailOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Col, Form, Input, Row } from 'antd';
import React from 'react';
import { history } from 'umi';
import './index.less';
export default function Login() {
  const onFinish = (values: any) => {
    register(values).then(() => {
      history.go(-1);
    });
  };
  return (
    <Row style={{ minHeight: '64vh' }}>
      <Col style={{ margin: 'auto' }}>
        <Col>
          <h1>用户注册 </h1>
        </Col>
        <Form onFinish={onFinish} className="login-form">
          <Form.Item
            name="name"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="请输入用户名"></Input>
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: '请输入用户密码' }]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              type="password"
              placeholder="请输入密码"
            />
          </Form.Item>
          <Form.Item name="email">
            <Input
              prefix={<MailOutlined />}
              placeholder="请输入邮件（非必需）"
              type="email"
            />
          </Form.Item>
          <Form.Item>
            <Button htmlType="submit" type="primary" style={{ width: '100%' }}>
              登陆
            </Button>
          </Form.Item>
        </Form>
      </Col>
    </Row>
  );
}
