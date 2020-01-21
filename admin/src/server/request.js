import axios from 'axios'
import custom from './custom'
const instance = axios.create({
    baseURL: 'http://127.0.0.1:8080',
    timeout:2500,
    // validateStatus:custom.err
});

instance.interceptors.request.use(...custom.req)

instance.interceptors.response.use(...custom.res)

export default instance