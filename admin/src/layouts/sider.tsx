import {
  HomeOutlined,
  LaptopOutlined,
  NotificationOutlined,
  PictureOutlined,
  SettingOutlined,
  UserOutlined,
} from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import React from 'react';
import { Link, useLocation } from 'umi';

const { SubMenu } = Menu;

export default function Sider(props: { collapsed: boolean | undefined }) {
  const location = useLocation();
  return (
    <Layout.Sider
      trigger={null}
      collapsible
      collapsed={props.collapsed}
      style={{ backgroundColor: '#fff' }}
    >
      <Menu
        mode="inline"
        selectedKeys={[location.pathname]}
        defaultOpenKeys={[location.pathname.split('/')[1]]}
        style={{ height: '100%', borderRight: 0 }}
      >
        <Menu.Item key="/" icon={<HomeOutlined />}>
          <Link to="">后台首页</Link>
        </Menu.Item>
        <SubMenu key="sub1" icon={<UserOutlined />} title="用户管理">
          <Menu.Item key="4">option4</Menu.Item>
        </SubMenu>
        <SubMenu key="post" icon={<LaptopOutlined />} title="文章管理">
          <Menu.Item key="/post">
            <Link to="/post">文章列表</Link>
          </Menu.Item>
          <Menu.Item key="/post/add">
            <Link to="/post/add">发布文章</Link>
          </Menu.Item>

          <Menu.Item key="/post/cate">
            <Link to="/post/cate">分类管理</Link>
          </Menu.Item>
          <Menu.Item key="/post/recomment">
            <Link to="/post/recomment">推荐管理</Link>
          </Menu.Item>
          <Menu.Item key="/post/tag">
            <Link to="/post/tag">标签统计</Link>
          </Menu.Item>
        </SubMenu>

        <SubMenu key="ad" icon={<PictureOutlined />} title="广告位管理">
          <Menu.Item key="/ads/groups">
            <Link to="/ads/groups">广告位管理</Link>
          </Menu.Item>
          <Menu.Item key="/ads/events">
            <Link to="/ads/events">事件管理</Link>
          </Menu.Item>
        </SubMenu>

        <SubMenu key="sub3" icon={<NotificationOutlined />} title="系统公告">
          <Menu.Item key="9">option9</Menu.Item>
          <Menu.Item key="10">option10</Menu.Item>
          <Menu.Item key="11">option11</Menu.Item>
          <Menu.Item key="12">option12</Menu.Item>
        </SubMenu>
        <Menu.Item key="/setting" icon={<SettingOutlined></SettingOutlined>}>
          <Link to="/setting">系统设置</Link>
        </Menu.Item>
      </Menu>
    </Layout.Sider>
  );
}
