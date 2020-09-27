import useRequest from '@/hooks/useRequest';
import { postRec } from '@/server/api/posts';
import { Button, Drawer, PageHeader } from 'antd';
import React, { useEffect, useState } from 'react';
import Post from '..';
export const PostsChoose = ({ id }: { id: number }) => {
  let { data: rec, load } = useRequest(() => postRec.detail(id), true);
  let [showSelect, setShowSelect] = useState(false);
  useEffect(() => {
    load();
  }, [id]);

  return (
    <div>
      <PageHeader title="选择文章" subTitle={'推荐位：首页' + id}></PageHeader>
      {JSON.stringify(rec)}
      <Button onClick={() => setShowSelect(true)}>选择</Button>
      <Drawer
        width={1200}
        visible={showSelect}
        onClose={() => setShowSelect(false)}
      >
        <Post
          showSelect={true}
          selectedKeys={[]}
          selectConfirm={(ids: number[]) => {
            setShowSelect(false);
            console.log(ids);

            postRec.update(id, { article_ids: ids.join(',') });
          }}
        ></Post>
      </Drawer>
    </div>
  );
};
