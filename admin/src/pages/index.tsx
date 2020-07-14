import useRequest from '@/hooks/useRequest';
import { posts } from '@/server/api/posts';
import { users } from '@/server/api/users';
import { Chart } from '@antv/g2';
import { Button } from 'antd';
import React, { useEffect } from 'react';
import { Link } from 'umi';
let chart: Chart;
let userChart: Chart;
export default () => {
  let { data: post, load: loadPost } = useRequest(() =>
    posts.total({ start: '2020-07-01 00:00:00' }),
  );
  let { data: user, load: loadUser } = useRequest(() =>
    users.total({ start: '2020-01-01 00:00:00' }),
  );
  const init = () => Promise.all([loadPost(), loadUser()]);

  useEffect(() => {
    init();
    chart = initChart('post-chart');
    userChart = initChart('user-chart');
  }, []);

  useEffect(() => {
    console.log(chart);
    if (chart) {
      chart.data(post.list || []);
      chart.render();
    }
    if (userChart) {
      userChart.data(user.list || []);
      userChart.render();
    }
  }, [post.list, user.list]);
  return (
    <div>
      <h1>文章数量:{post.total || 0}</h1>
      <div id="post-chart"></div>
      <h1>用户数量:{user.total || 0}</h1>
      <div id="user-chart"></div>

      <Button type="primary">
        <Link to="/setting">hello world!</Link>
      </Button>
    </div>
  );
};

const initChart = (container: string): Chart => {
  let chart = new Chart({
    container,
    autoFit: false,
    height: 300,
    width: 500,
    padding: [40, 40, 40, 40],
  });
  chart.data([]);
  chart.scale({
    date: {
      nice: true,
    },
    count: {
      nice: true,
      alias: '当天总数',
    },
    offset: {
      alias: '增长',
    },
  });

  chart.tooltip({
    showCrosshairs: true, // 展示 Tooltip 辅助线
    shared: true,
  });

  chart
    .line()
    .position('date*offset')
    .color('#eee')
    .shape('dash');

  chart
    .line()
    .position('date*count')
    .color('blue');

  // chart.interval().position('date*count');

  // chart
  //   .line()
  //   .position('date*count')
  //   .label('count');
  chart.point().position('date*count');
  return chart;
};
