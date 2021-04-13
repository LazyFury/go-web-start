import config from '@/utils/config';
import { message } from 'antd';
import axios, { AxiosRequestConfig } from 'axios';
import { history } from 'umi';

export const http = axios.create({
    baseURL: config.baseURL
});

http.interceptors.request.use((config: AxiosRequestConfig) => {
    let Authorization = window.localStorage.getItem('token') || '';
    let { headers = {} } = config;
    let _config = {
        ...config,
        headers: {
            ...headers,
            Authorization
        },
        withCredentials: true
    };
    return _config;
});

http.interceptors.response.use(interceptorsResponse, onrejectionhandled);
function onrejectionhandled(err: any) {
    if (err.response) {
        let res = err?.response?.data;
        message.error(res?.msg || '请求失败1');
        handleErrCode(res?.code, res);
        return;
    }
    message.error('请求失败');
}
function interceptorsResponse(res: any) {
    // console.log(res);
    let data: any = res.data;
    let success: boolean = data.code >= 200 && data.code < 300;
    let msg: string = data.message || '';
    const code: number = data.code;
    // let result = data;
    // 成功;
    if (success) {
        let ignore = ['请求成功'];
        if (!ignore.includes(msg)) {
            message.success({ content: msg });
        }
        return res;
    } else {
        message.error({ content: msg });
        handleErrCode(code, data);
    }
    return Promise.reject({ err: res, text: '请求失败' });
}

function handleErrCode(code: number, data: any) {
    switch (code) {
        case -100:
            history.push('/login');
            break;
    }
}
