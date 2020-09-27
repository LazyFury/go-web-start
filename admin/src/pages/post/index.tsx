import { useDataList } from '@/hooks/useDataList';
import useRequest from '@/hooks/useRequest';
import { postCates, posts } from '@/server/api/posts';
import config from '@/utils/config';
import {
  CheckOutlined,
  DeleteOutlined,
  EditOutlined,
  PaperClipOutlined,
  SyncOutlined,
} from '@ant-design/icons';
import { Button, Modal, PageHeader, Select, Table, Tooltip } from 'antd';
import React, { useEffect, useState } from 'react';
import { Link } from 'umi';
import './index.less';

const { Option } = Select;
let { confirm } = Modal;
let resetTableData: () => Promise<void>; //在其他组件中使用重置列表
let setCate_id: React.Dispatch<React.SetStateAction<string>>;

let numbers: number[] = [];

var 补集 = (a: number[], b: number[]) =>
  a
    .filter(function(v) {
      return !(b.indexOf(v) > -1);
    })
    .concat(
      b.filter(function(v) {
        return !(a.indexOf(v) > -1);
      }),
    );

const Post = (props: {
  showSelect?: boolean;
  selectedKeys?: number[];
  selectConfirm?: (keys: number[]) => void;
}) => {
  /**
   * @init
   */
  let [cid, setCid] = useState('');
  setCate_id = setCid;

  let [selectedRowKeys_, setSelectedRowKeys] = useState(numbers);

  // 获取文章列表
  let { data, load, loading, reset } = useDataList(
    page => posts.list({ page, cid }),
    false,
  );

  let { data: cates } = useRequest(() => postCates.list(), true);

  resetTableData = load;

  /**
   * @event
   */

  useEffect(() => {
    reset();
  }, [cid]);

  useEffect(() => {
    setSelectedRowKeys(props.selectedKeys || []);
  }, [props.selectedKeys]);

  /**
   * @methods
   */
  const cateSelectChange = (value: string) => {
    setCid(value);
  };

  /**
   * @render
   */
  return (
    <div>
      <PageHeader
        style={{ padding: '20rpx 0' }}
        // onBack={() => null}
        title="文章管理"
        subTitle={`共${data.count}篇文章，${data.page_count}页，当前${data.page_now}页`}
      />

      <div className="page-main">
        <div className="action-bar" style={{ margin: '10px 0' }}>
          {props.showSelect && (
            <Button
              type="primary"
              onClick={() => {
                props.selectConfirm && props.selectConfirm(selectedRowKeys_);
                setSelectedRowKeys([]);
              }}
            >
              <CheckOutlined />
              确定选择
            </Button>
          )}

          {/* 选择分类 */}
          <Select
            allowClear
            placeholder="请选择文章分类"
            value={cid || '全部分类'}
            onChange={cateSelectChange}
          >
            {/* cid 使用usestate， 默认值为0而不是空 */}
            <Option key={0} value={0}>
              全部分类
            </Option>
            {cates && cates.length > 0
              ? cates.map((x: any) => {
                  return (
                    <Option key={x.id} value={x.id}>
                      {x.name}
                    </Option>
                  );
                })
              : null}
          </Select>
          {/* 刷新列表 */}
          <Button onClick={() => load()}>
            <SyncOutlined></SyncOutlined>
            <span>刷新</span>
          </Button>
        </div>

        <Table
          columns={columns}
          dataSource={data.list}
          size={'large'}
          bordered={true}
          rowKey={record => record.id}
          loading={loading}
          pagination={{
            position: ['bottomLeft'],
            current: data.page_now,
            total: data.count,
            showSizeChanger: false,
            onChange: load,
          }}
          rowSelection={
            (props.showSelect && {
              type: 'checkbox',
              selectedRowKeys: selectedRowKeys_,
              onSelectAll: (selected: boolean, selectedRows, changeRows) => {
                let ids = [];
                if (!selected) {
                  ids = changeRows.map(x => x.id);
                  setSelectedRowKeys(补集(selectedRowKeys_, ids));
                }
              },

              onSelect: (item, selected: boolean, selectedRows: any[]) => {
                console.log(item, selected);
                let key = item.id;
                let selectedKeys = selectedRowKeys_;
                if (!selected) {
                  let index = selectedKeys.indexOf(key);
                  if (index > -1) {
                    selectedKeys.splice(index, 1);
                  }
                }
                setSelectedRowKeys(selectedKeys);
              },
              onChange: (selectedRowKeys, selectedRows) => {
                let keys = selectedRows.map(x => x.id);
                setSelectedRowKeys([
                  ...new Set([...selectedRowKeys_, ...keys]),
                ]);
              },
            }) ||
            undefined
          }
        ></Table>
      </div>
    </div>
  );
};

// 表格列配置
const columns = [
  { title: 'ID', key: 'id', dataIndex: 'id' },
  { title: '文章标题', key: 'title', dataIndex: 'title', render: title },
  // {
  //   title: '简介',
  //   key: 'desc',
  //   dataIndex: 'desc',
  //   render(desc: any) {
  //     return <div>{desc || '暂无内容....'}</div>;
  //   },
  // },
  {
    title: '分类',
    key: 'cate',
    dataIndex: 'cate_name',
    render: filterCate,
  },
  { title: '作者', key: 'author', dataIndex: 'author' },
  { title: '文章标签', key: 'tag', dataIndex: 'tag' },
  { title: '创建时间', key: 'created_at', dataIndex: 'created_at' },
  { title: '更新时间', key: 'updated_at', dataIndex: 'updated_at' },
  {
    title: '操作',
    key: 'action',
    dataIndex: 'id',
    render: (id: any) => {
      return (
        <div style={{ width: 120 }}>
          {edit(id)}
          {del(id)}
        </div>
      );
    },
  },
];

// 根据分类选择
function filterCate(_: any, data: any) {
  return (
    <div>
      {data.cate_name}
      <a onClick={() => setCate_id(data.cate_id)}>
        <PaperClipOutlined></PaperClipOutlined>
      </a>
    </div>
  );
}

function del(id: number) {
  return (
    <Tooltip title="删除文章">
      <a
        onClick={() => confirmDel(id)}
        style={{ marginLeft: '4px', color: '#d33', padding: '0 10px' }}
      >
        <span>删除</span>
      </a>
    </Tooltip>
  );
}

function edit(id: number) {
  return (
    <Tooltip title="编辑文章">
      <Link to={'/post/add?type=edit&id=' + id} style={{ marginLeft: '4px' }}>
        <EditOutlined></EditOutlined>
        <span>编辑</span>
      </Link>
    </Tooltip>
  );
}

function title(article: string, data: any) {
  return (
    <div>
      <span> {article}</span>
      <Tooltip title="预览文章">
        <a
          href={config.previewUrl + '/posts/' + data.id}
          target="_blank"
          style={{ marginLeft: '4px' }}
        >
          <PaperClipOutlined></PaperClipOutlined>
        </a>
      </Tooltip>
    </div>
  );
}

function confirmDel(id: number) {
  return confirm({
    title: '确定要删除文章？无法撤回',
    icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
    okText: '确定删除',
    okType: 'danger',
    onOk() {
      posts.del(id).then(res => {
        resetTableData();
      });
    },
  });
}

export default Post;
