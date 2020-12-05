import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { postTags } from '@/server/api/posts';
import { randomColor } from '@/utils/utils';
import { Tag } from 'antd';
import React from 'react';

export default () => {
  let { data: tags, load: loadTags } = useRequest(postTags.list, true);
  return (
    <PageMain
      title="标签统计"
      subTitle={`这个接口设计有性能问题，计数使用的程序内遍历，需要重新设计数据表结构,或者改为惰性统计count存到数据库`}
    >
      <div>
        {tags instanceof Array &&
          tags.map(tag => {
            return (
              <Tag
                key={tag.id}
                color={randomColor()}
                style={{ marginBottom: '10px' }}
              >
                {tag.val || '~'}({tag.count})
              </Tag>
            );
          })}
      </div>
    </PageMain>
  );
};
