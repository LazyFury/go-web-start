import {
  FullscreenOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
} from '@ant-design/icons';
import { Col, Layout, Row } from 'antd';
import React, { useState } from 'react';
import './index.less';

const iconStyle = { fontSize: '18px' };
export default function Header(props: {
  collapsed: Boolean;
  toggle: (event: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
}) {
  const [fullScreen, SetFullScreen] = useState(false);

  const fullScreenBtn = () => {
    return React.createElement(
      fullScreen ? FullscreenOutlined : FullscreenOutlined,
      {
        onClick: () => {
          const body = document.querySelector('body');
          if (body) {
            if (!fullScreen) {
              body.requestFullscreen({}).then(() => {
                SetFullScreen(true);
              });
            } else {
              document.exitFullscreen().then(() => {
                SetFullScreen(false);
              });
            }
          }
        },
        style: { color: '#fff', ...iconStyle },
      },
    );
  };

  return (
    <Layout.Header>
      <Row style={{ alignItems: 'center' }}>
        <div className="logo"></div>
        {React.createElement(
          props.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
          {
            className: 'trigger',
            onClick: props.toggle,
            style: { color: '#fff', ...iconStyle },
          },
        )}
        <Col flex="auto"></Col>

        {fullScreenBtn()}
      </Row>
    </Layout.Header>
  );
}
