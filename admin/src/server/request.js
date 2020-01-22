import axios from 'axios'
import custom from './custom'
import config from '../config'
const instance = axios.create({
    baseURL: config.baseURL,
    timeout:2500,
    // validateStatus:custom.err
});

instance.interceptors.request.use(...custom.req)

instance.interceptors.response.use(...custom.res)

export default instance