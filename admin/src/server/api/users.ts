import { history } from 'umi';
import { http } from '../request';
import { install } from './easy_install';

export const users = {
    ...install('users'),
    profile: () => http.get('/user-profile')
};

export const login = (data: any) => {
    return http.post('/login', data).then(res => {
        window.localStorage.setItem('token', res.data.data);
        history.go(-1);
        return res;
    });
};

export const register = (data: any) => {
    return http.post('/reg', data).then(res => {});
};
