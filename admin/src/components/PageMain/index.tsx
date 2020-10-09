import { Card, Divider, PageHeader } from 'antd';
import React from 'react';

export default function(props: {
  children: React.ReactNode;
  title: string;
  subTitle: string;
}) {
  return (
    <div>
      <Card bordered>
        <PageHeader
          style={{ padding: '0' }}
          // onBack={() => null}
          title={props.title}
          subTitle={props.subTitle}
        />
      </Card>

      <Divider dashed style={{ margin: '10px 0' }} />

      <Card bordered>{props.children}</Card>
    </div>
  );
}
