import {
    HomeOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    PlusOutlined,
    UserOutlined,
} from "@ant-design/icons";
import { Layout, Menu, Button, theme, FloatButton } from "antd";
import React from "react";
import { useState } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";

const { Header, Sider, Content } = Layout;

export default function Root() {
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  return (
    <React.Fragment>
      <Layout
        style={{
          height: "100vh",
          margin: 0,
        }}
      >
        <Sider trigger={null} collapsible collapsed={collapsed}>
          <div className="demo-logo-vertical" />
          <Menu
            theme="dark"
            mode="inline"
            defaultSelectedKeys={["1"]}
            items={[
              {
                key: "1",
                icon: <Link to={"/"}><HomeOutlined /></Link>,
                label: "Home",
              },
              {
                key: "2",
                icon: <Link to={"/employees"}><UserOutlined /></Link>,
                label: "Employees",
              },
            ]}
          />
        </Sider>
        <Layout>
          <Header style={{ padding: 0, background: colorBgContainer }}>
            <Button
              type="text"
              icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              onClick={() => setCollapsed(!collapsed)}
              style={{
                fontSize: "16px",
                width: 64,
                height: 64,
              }}
            />
          </Header>
          <Content
            style={{
              margin: "24px 16px",
              padding: 24,
              minHeight: 280,
              background: colorBgContainer,
            }}
          >
            <Outlet />
          </Content>
        </Layout>
      </Layout>
      <FloatButton
        shape="circle"
        type="primary"
        style={{ right: 36 }}
        icon={<PlusOutlined />}
        onClick={() => navigate("/employees/add")}
      />
    </React.Fragment>
  );
}
