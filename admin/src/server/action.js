import { message } from "ant-design-vue";
import Route from '../router'

const login = () => {
  message.error("请先登陆");
  setTimeout(() => {
    Route.push('/login/login')
  }, 100);
  return Promise.reject(res)
};

export default {
  login
};
