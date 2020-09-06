import config from '@/utils/config';
import { message } from 'antd';
import axios, { AxiosRequestConfig } from 'axios';
import { history } from 'umi';

export const http = axios.create({
  baseURL: config.baseURL,
});

http.interceptors.request.use((config: AxiosRequestConfig) => {
  let Authorization = window.localStorage.getItem('token') || '';
  let { headers = {} } = config;
  let _config = {
    ...config,
    headers: {
      ...headers,
      Authorization,
    },
    withCredentials: true,
  };
  return _config;
});

http.interceptors.response.use(interceptorsResponse);

function interceptorsResponse(res: any) {
  // console.log(res);
  let data: any = res.data;
  let success: boolean = data.code == 200;
  let msg: string = (data && data.msg) || '';
  const code: number = data.code;
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
    handleErrCode(code);
  }
  return Promise.reject({ err: res, text: '请求失败' });
}

function handleErrCode(code: number) {
  switch (code) {
    case -101:
      history.push('/login');
      break;
  }
}
