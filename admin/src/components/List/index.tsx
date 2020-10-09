import {
  DownloadOutlined,
  SettingOutlined,
  SyncOutlined,
  VerticalAlignMiddleOutlined,
} from '@ant-design/icons/';
import {
  Button,
  Checkbox,
  Col,
  Divider,
  Popover,
  Row,
  Space,
  Tooltip,
} from 'antd';
import { SizeType } from 'antd/es/config-provider/SizeContext';
import { CheckboxOptionType } from 'antd/lib/checkbox';
import { CheckboxValueType } from 'antd/lib/checkbox/Group';
import Table, { ColumnsType, TableProps } from 'antd/lib/table';
import React, { useEffect, useState } from 'react';
import './index.less';

export default function(props: {
  children?: React.ReactNode;
  leftActions?: React.ReactNode[];
  onRefresh?: () => void;
  loading?: boolean;
  table: TableProps<any>;
}) {
  let leftActions = props.leftActions || [];

  let columns: ColumnsType = (props.table || {}).columns || [];

  const defaultTableSize: SizeType[] = ['small', 'middle', 'large'];
  const [tableSize, setTableSize] = useState(1);

  let defaultPlainOptions: string[] = [];
  const [plainOptions, setplainOptions] = useState(defaultPlainOptions);

  useEffect(() => {
    setplainOptions(columns.map(x => x.key + ''));
  }, []);

  const IconBtn = (props: {
    title: string;
    onClick: () => void;
    icon: React.ReactNode;
  }) => {
    return (
      <Tooltip title={props.title} placement="top">
        <Button
          type="text"
          style={{ fontSize: '18px', color: '#666' }}
          onClick={props.onClick}
        >
          <Row style={{ alignItems: 'center' }}>
            {props.icon}
            <span
              style={{
                fontSize: '12px',
                marginLeft: '6px',
              }}
            >
              {props.title}
            </span>
          </Row>
        </Button>
      </Tooltip>
    );
  };

  const moreActions = () => {
    return (
      <Row>
        {/* <Space> */}
        <IconBtn
          title="刷新"
          onClick={() => {
            props.onRefresh instanceof Function && props.onRefresh();
          }}
          icon={<SyncOutlined />}
        />

        <Popover content={AlignSet()} trigger="click" placement="bottom">
          <IconBtn
            title="间距设置"
            onClick={() => {}}
            icon={<VerticalAlignMiddleOutlined />}
          />
        </Popover>

        <Popover content={ItemSet()} trigger="click" placement="bottom">
          <IconBtn
            title=" 列设置"
            onClick={() => {}}
            icon={<SettingOutlined />}
          />
        </Popover>

        <IconBtn
          title="导出数据"
          onClick={() => {}}
          icon={<DownloadOutlined />}
        />
        {/* </Space> */}
      </Row>
    );
  };

  const ItemSet = () => {
    const options = columns.map(x => {
      let value: CheckboxOptionType = {
        label: x.title + '',
        value: x.key || '',
      };
      return value;
    });

    const defaultValue = options.map(x => x.value + '');
    return (
      <Checkbox.Group
        style={{ display: 'flex', flexDirection: 'column' }}
        options={options}
        defaultValue={defaultValue}
        onChange={onItemSetChange}
      />
    );
  };

  const onItemSetChange = (e: CheckboxValueType[]) => {
    setplainOptions(e.map(x => String(x)));
  };

  const AlignSet = () => {
    return (
      <Col>
        {defaultTableSize.map((x, i) => {
          return (
            <div key={x}>
              <a
                style={{ color: tableSize !== i ? '#666' : '' }}
                onClick={() => setTableSize(i)}
              >
                {x}
              </a>
              {i < defaultTableSize.length - 1 && (
                <Divider style={{ margin: '4px 0' }} />
              )}
            </div>
          );
        })}
      </Col>
    );
  };

  return (
    <div className="table-list">
      <Row
        className="table-action-bar"
        style={{ margin: '10px 0', alignItems: 'center' }}
      >
        {/* 左边按钮 */}
        <Col>
          <Space>{leftActions}</Space>
        </Col>
        {/* 自动撑开 */}
        <Col flex="auto"></Col>
        {/* 右边按钮 */}
        {moreActions()}
      </Row>

      <Divider></Divider>

      {/* 内容 */}
      <Table
        {...props.table}
        size={defaultTableSize[tableSize]}
        columns={columns.filter(x => plainOptions.includes(x.key + ''))}
      />
      {/* {props.children} */}
    </div>
  );
}
