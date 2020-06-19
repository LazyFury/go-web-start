import { Breadcrumb, Layout as ALayout } from 'antd';
import React, { useState } from 'react';
import Header from './header';
import Sider from './sider';
const Layout = (props: { children: React.ReactNode }) => {
  let [collapsed, setCollapsed] = useState(false);

  const toggle = () => {
    setCollapsed(!collapsed);
  };
  return (
    <ALayout style={{ minHeight: '100vh' }}>
      {/* header  */}
      <Header collapsed={collapsed} toggle={toggle}></Header>
      <ALayout>
        {/* sider */}
        <Sider collapsed={collapsed}></Sider>
        <ALayout style={{ padding: '20px' }}>
          <Breadcrumb></Breadcrumb>
          {/* content */}
          <ALayout.Content className="fff" style={{ padding: '20px' }}>
            {props.children}
          </ALayout.Content>
        </ALayout>
      </ALayout>
    </ALayout>
  );
};

export default Layout;
