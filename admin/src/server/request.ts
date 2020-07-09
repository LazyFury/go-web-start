import config from '@/utils/config';
import { message } from 'antd';
import axios, { AxiosRequestConfig } from 'axios';

export const http = axios.create({ baseURL: config.baseURL });

http.interceptors.request.use((config: AxiosRequestConfig) => config);
http.interceptors.response.use(interceptorsResponse);

function interceptorsResponse(res: any) {
  console.log(res);
  let data = res.data;
  let success = data.code == 200;
  let msg = (data && data.msg) || '';
  let result = data;
  // 成功;
  if (success) {
    let ignore = ['请求成功'];
    if (!ignore.includes(msg)) {
      message.success({ content: msg });
    }
    return result;
  } else {
    message.error({ content: msg });
  }
  return Promise.reject({ err: res, text: '请求失败' });
}
