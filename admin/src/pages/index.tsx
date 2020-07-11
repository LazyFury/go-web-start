import useRequest from '@/hooks/useRequest';
import { posts } from '@/server/api/posts';
import { users } from '@/server/api/users';
import { Button } from 'antd';
import React, { useEffect } from 'react';
import { Link } from 'umi';

export default () => {
  let { data: post, load: loadPost } = useRequest(() =>
    posts.total({ start: '2020-07-01 00:00:00' }),
  );
  let { data: user, load: loadUser } = useRequest(() =>
    users.total({ start: '2020-06-01 00:00:00' }),
  );
  const init = () => Promise.all([loadPost(), loadUser()]);

  useEffect(() => {
    init();
  }, []);
  return (
    <div>
      <h1>文章数量:{post.total || 0}</h1>
      <h1>用户数量:{user.total || 0}</h1>
      <Button type="primary">
        <Link to="/setting">hello world!</Link>
      </Button>
    </div>
  );
};
