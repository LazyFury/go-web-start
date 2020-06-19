import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { Layout } from 'antd';
import React from 'react';
import './index.less';

export default function Header(props: {
  collapsed: Boolean;
  toggle: (event: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
}) {
  return (
    <Layout.Header>
      <div className="logo"></div>
      {React.createElement(
        props.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
        {
          className: 'trigger',
          onClick: props.toggle,
          style: { color: '#fff' },
        },
      )}
    </Layout.Header>
  );
}
