import action from './action'
import {message} from 'ant-design-vue'

let ignoreCode = ["请求成功"]
// 错误码拦截
let errCode = {
    '1':(res)=>{
        console.info(res.msg,"Success")
        if(!ignoreCode.includes(res.msg)){
            message.success(res.msg)
        }
        return res
    },
    '-1':(res)=>{
        message.error(res.msg)
        return Promise.reject(res)
    },
    '-101':action.login,
    '-102':action.login,
}

export default errCode