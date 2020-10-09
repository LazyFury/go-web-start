import List from '@/components/List';
import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { postRec } from '@/server/api/posts';
import { randomColor } from '@/utils/utils';
import { Button, Drawer, Tag } from 'antd';
import { ColumnsType } from 'antd/lib/table';
import React, { useEffect, useState } from 'react';
import Post from '..';

let listNumber: number[];

export const PostsChoose = ({ id }: { id: number }) => {
  let { data: rec, load, loading } = useRequest(() => postRec.detail(id), true);
  let [showSelect, setShowSelect] = useState(false);

  let [selectedKeys, setSelectedKeys] = useState(listNumber);

  useEffect(() => {
    load();
  }, [id]);

  let getIds = (data: { list: any[] }) => {
    let arr: any[] = data.list || [];
    return arr.map((x: { id: number }) => x.id);
  };

  const columns: ColumnsType<any> = [
    { title: 'ID', key: 'id', dataIndex: 'id' },
    { title: '文章标题', key: 'title', dataIndex: 'title' },
    {
      title: '分类',
      key: 'cate',
      dataIndex: 'cate_name',
    },
    { title: '作者', key: 'author', dataIndex: 'author' },
    {
      title: '文章标签',
      key: 'tag',
      dataIndex: 'tag',
      render: (tag: any[]) =>
        tag.slice(0, 3).map((text: string) => (
          <Tag key={text} style={{ marginBottom: '4px' }} color={randomColor()}>
            {text}
          </Tag>
        )),
    },
    // { title: '创建时间', key: 'created_at', dataIndex: 'created_at' },
    // { title: '更新时间', key: 'updated_at', dataIndex: 'updated_at' },
  ];
  return (
    <PageMain title="选择文章" subTitle={'推荐位：首页' + id}>
      <List
        onRefresh={load}
        loading={loading}
        leftActions={[
          <Button
            onClick={() => {
              setSelectedKeys(getIds(rec));
              setShowSelect(true);
            }}
            type="primary"
          >
            选择
          </Button>,
          <Button
            onClick={() => {
              postRec.update(id, { article_ids: '0' }).then(() => load());
            }}
          >
            清空
          </Button>,
        ]}
        table={{
          columns,
          dataSource: rec.list instanceof Array ? rec.list : [],
          loading,
        }}
      ></List>

      <Drawer
        width={1200}
        visible={showSelect}
        onClose={() => setShowSelect(false)}
      >
        <Post
          showSelect={true}
          selectedKeys={selectedKeys}
          selectConfirm={(ids: number[]) => {
            setShowSelect(false);
            console.log(ids);

            postRec.update(id, { article_ids: ids.join(',') }).then(() => {
              load();
            });
          }}
        ></Post>
      </Drawer>
    </PageMain>
  );
};
