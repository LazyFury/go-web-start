
import http from './request'


const get = (url,option) =>{
    return http.get(url,{...option,params:option.data})
}

const post = (url,option) =>{
    return http.post(url,option.data,{...option})
}

const api = {
    login(data){
        return get('/admin/login',{data}).then(res=>{
            if(res.code==1){
                window.localStorage.setItem("token",res.data)
            }
            return res
        })
    },
    adminHome(data){
        return get('/admin',{data})
    },
    // 用户接口
    user:{
        list(data){
            return get('/admin/user/list',{data})
        },
        update(data){
            return post('/admin/user/updateUser',{data})
        },
        frozen(data){
            return post('/admin/user/frozen',{data})
        },
        del(data){
            return post('/admin/user/delUser',{data})
        }
    }
}

export {
    api,http
}