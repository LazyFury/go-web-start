
/**
 * 从url中获取参数
 * @param url 
 */
 function GetParam(url=""){
    if (url == ""){
        url = window.location.href
    }
    let paramstr = url.split("?")[1]
    let arr = []
    let param = {}
    if(paramstr){
        arr = paramstr.split('&')
    }
    if(arr.length>0){
        arr.forEach(x=>{
            let array_ = x.split('=')
            param[array_[0]] = param[array_[1]]
        })
    }

    return param
}

const util = {
    GetParam
}

export default util