import { login } from '@/server/api/users';
import { Button, Form, Input } from 'antd';
import React from 'react';

export default function Login() {
  const onFinish = (values: any) => {
    login(values);
  };
  return (
    <div>
      <Form onFinish={onFinish}>
        <Form.Item name="username" label="Username">
          <Input></Input>
        </Form.Item>
        <Form.Item name="password" label="Password">
          <Input></Input>
        </Form.Item>
        <Form.Item>
          <Button htmlType="submit">登陆</Button>
        </Form.Item>
      </Form>
    </div>
  );
}
