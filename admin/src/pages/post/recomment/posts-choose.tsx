import useRequest from '@/hooks/useRequest';
import { postRec } from '@/server/api/posts';
import { Button, Drawer, PageHeader } from 'antd';
import React, { useEffect, useState } from 'react';
import Post from '..';

let listNumber: number[];

export const PostsChoose = ({ id }: { id: number }) => {
  let { data: rec, load } = useRequest(() => postRec.detail(id), true);
  let [showSelect, setShowSelect] = useState(false);

  let [selectedKeys, setSelectedKeys] = useState(listNumber);

  useEffect(() => {
    load();
  }, [id]);

  let getIds = (data: { list: any[] }) => {
    let arr: any[] = data.list || [];
    return arr.map((x: { id: number }) => x.id);
  };

  return (
    <div>
      <PageHeader title="选择文章" subTitle={'推荐位：首页' + id}></PageHeader>
      <Button
        onClick={() => {
          setSelectedKeys(getIds(rec));
          setShowSelect(true);
        }}
      >
        选择
      </Button>
      <Button
        onClick={() => {
          postRec.update(id, { article_ids: '0' }).then(() => load());
        }}
      >
        清空
      </Button>
      {JSON.stringify(rec.list)}

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
    </div>
  );
};
