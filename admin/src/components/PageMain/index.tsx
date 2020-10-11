import { Card, Col, Divider, Typography } from 'antd';
import React from 'react';
const { Text, Title } = Typography;

export default function(props: {
  children: React.ReactNode;
  title: string;
  subTitle: string;
}) {
  return (
    <div>
      <Card bordered>
        <Col>
          <Title level={3}>{props.title}</Title>
          <Text type="secondary">{props.subTitle}</Text>
        </Col>
      </Card>

      <Divider dashed style={{ margin: '10px 0' }} />

      <Card bordered>{props.children}</Card>
    </div>
  );
}
