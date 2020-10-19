import { Button, Result } from 'antd';
import React from 'react';
import { history } from 'umi';

export default function NotFund() {
  return (
    <div>
      <Result
        status="404"
        title="找不到页面，请稍后重试"
        subTitle="404 Not Fund"
        extra={
          <Button type="primary" onClick={() => history.go(-2)}>
            返回上一页
          </Button>
        }
      ></Result>
    </div>
  );
}
