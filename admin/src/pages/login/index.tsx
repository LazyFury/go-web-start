import { login } from '@/server/api/users';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Col, Form, Input, Row } from 'antd';
import React from 'react';
import { history } from 'umi';
import './index.less';
export default function Login() {
    const onFinish = (values: any) => {
        login(values).then(() => {
            history.go(-1);
        });
    };
    return (
        <Row style={{ minHeight: '64vh' }}>
            <Col style={{ margin: 'auto' }}>
                <Col>
                    <h1>用户登陆 </h1>
                </Col>
                <Form onFinish={onFinish} className="login-form">
                    <Form.Item name="name" rules={[{ required: true, message: '请输入用户名' }]}>
                        <Input prefix={<UserOutlined />} placeholder="请输入用户名"></Input>
                    </Form.Item>
                    <Form.Item name="password" rules={[{ required: true, message: '请输入用户密码' }]}>
                        <Input.Password prefix={<LockOutlined />} type="password" placeholder="请输入密码" />
                    </Form.Item>
                    <Form.Item>
                        <Row>
                            <Col flex={1}>
                                没有账号？
                                <a
                                    onClick={() => {
                                        history.push('/login/register');
                                    }}
                                >
                                    立即注册
                                </a>
                            </Col>
                            <a
                                onClick={() => {
                                    history.push('/login/forgot');
                                }}
                            >
                                忘记密码
                            </a>
                        </Row>
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
