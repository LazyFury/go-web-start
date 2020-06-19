import { Button, PageHeader } from 'antd';
import React from 'react';
import { Link } from 'umi';
import styles from './index.less';

export default () => {
  return (
    <div>
      <PageHeader
        className="site-page-header fff"
        onBack={() => null}
        title="Title"
        subTitle="This is a subtitle"
      />
      <h1 className={styles.title}>Page index</h1>
      <Button type="primary">
        <Link to="/setting">hello world!</Link>
      </Button>
    </div>
  );
};
