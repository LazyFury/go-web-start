import useRequest from '@/hooks/useRequest';
import { postTags } from '@/server/api/posts';
import { randomColor } from '@/utils/utils';
import { PageHeader, Tag } from 'antd';
import React from 'react';

export default () => {
  let { data: tags, load: loadTags } = useRequest(postTags.list, true);
  return (
    <div>
      <PageHeader
        style={{ padding: '20rpx 0' }}
        // onBack={() => null}
        title="标签统计"
        subTitle={``}
      />
      <div className="page-main">
        {tags instanceof Array &&
          tags.map(tag => {
            return (
              <Tag key={tag.id} color={randomColor()}>
                {tag.val || '~'}({tag.count})
              </Tag>
            );
          })}
      </div>
    </div>
  );
};
