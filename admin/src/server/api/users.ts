import { history } from 'umi';
import { http } from '../request';
import { install } from './easy_install';

export const users = {
  ...install('users'),
};

export const login = (data: any) => {
  return http.post('/login', data).then(res => {
    window.localStorage.setItem('token', res.data);
    history.go(-1);
    return res;
  });
};

export const register = (data: any) => {
  return http.post('/login/reg', data).then(res => {});
};
