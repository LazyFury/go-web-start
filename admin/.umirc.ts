import { defineConfig } from 'umi';

export default defineConfig({
  title: '后台管理',
  nodeModulesTransform: {
    type: 'none',
    exclude: [],
  },
  // ssr: {
  //   // mode: 'stream',
  // },
  // routes: [{ path: '/', component: '@/pages/index' }],
});
