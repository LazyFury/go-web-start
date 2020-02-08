
/**
 * 从url中获取参数
 * @param url 
 */
 function GetParam(url=""){
    if (url == ""){
        url = window.location.href
    }
    let urls = url.split("?")
    let paramstr = urls[urls.length-1] //
    // 存储分隔后的字符串
    let arr = []
    // 存储参数对象
    let param = {}
    // 分隔参数
    if(paramstr){
        arr = paramstr.split('&')
    }
    // 赋值参数
    if(arr.length>0){
        arr.forEach(x=>{
            let array_ = x.split('=')
            param[array_[0]] = array_[1]
        })
    }
    console.info("Get created Param:",param,"in url(",url,")")
    return param
}


export default {
    GetParam
}