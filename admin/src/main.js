// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
Vue.config.productionTip = false
// 工具
import util from './util/util'
// api 
import { api } from './server/api'
import './util/prototype'
import custom_plugin from './util/plugin'
import store from './util/store';

// import './util/test'

Vue.prototype.$util = util
Vue.prototype.api = api
Vue.prototype.$store = store
Vue.prototype.$isDev = false
Vue.use(custom_plugin)
Vue.use(Antd)
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
