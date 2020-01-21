import {notification} from 'ant-design-vue'
import errCode from './errCode'
// 前置请求拦截
const reqHandle = (config) => {
    console.log("请求配置：",config)
    let {headers} = config
    headers['token'] = window.localStorage.getItem("token")

    config = {...config,headers}
    return config
};

const reqErrHandle = (err) => {
    return Promise.reject(err)
};


// 返回结果拦截
const resHandle = (response) => {
    let {data:res} = response
    console.log("请求结果:",res)
    res.code += "" //转换文本类型 
    // console.log(res.code in errCode,res.code,errCode)
    // 错误码处理  正常的逻辑 错误码为1的时候处理
    if (res.code in errCode) {
        if(errCode[res.code] instanceof Function){
            return errCode[res.code](res)
        }
    }
    
    return res
};

const resErrHandle = (err) => {
    console.error('请求失败',err)
    notification.error({
        message: '请求失败',
        description:'-'
    })
    return Promise.reject(err)
};

const errHandle = (status) => {
    return status!=200
};


let custom = {
    req:[reqHandle,reqErrHandle],
    res:[resHandle,resErrHandle],
    err:errHandle
  }
export default custom;
