import { Layout as ALayout } from 'antd';
import React, { useState } from 'react';
import { history } from 'umi';
import Header from './header';
import Sider from './sider';

let noLayout = ['/login', '/login/forgot', '/login/register'];

const Layout = (props: {
  location: any;
  route: any;
  children: React.ReactNode;
}) => {
  let [collapsed, setCollapsed] = useState(false);

  const toggle = () => {
    setCollapsed(!collapsed);
  };

  // 独立于layout
  if (noLayout.includes(props.location.pathname)) {
    return <div>{props.children}</div>;
  }

  console.log(props);

  // 完美匹配时找不到页面，不支持 [/posts/:id] 式的路由,约定式路由的404自动配置不生效 自动匹配到根目录/
  let { pathname } = props.location;
  let { routes = [] } = props.route;
  if (!routes.map((x: { path: string }) => x.path).includes(pathname)) {
    history.push('/error/404');
  }

  // debugger;
  return (
    <ALayout style={{ minHeight: '100vh' }}>
      {/* header  */}
      <Header collapsed={collapsed} toggle={toggle}></Header>
      <ALayout>
        {/* sider */}
        <Sider collapsed={collapsed}></Sider>
        <ALayout style={{ padding: '20px' }}>
          {/* <Breadcrumb></Breadcrumb> */}
          {/* content */}
          <ALayout.Content className="">{props.children}</ALayout.Content>
        </ALayout>
      </ALayout>
    </ALayout>
  );
};

export default Layout;
