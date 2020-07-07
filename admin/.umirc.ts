import { defineConfig } from 'umi';

export default defineConfig({
  title: 'hello umi admin',
  nodeModulesTransform: {
    type: 'none',
    exclude: [],
  },
  // routes: [{ path: '/', component: '@/pages/index' }],
});
