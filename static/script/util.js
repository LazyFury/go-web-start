function curlPost(url,data,callback){
    return fetch(url,{
        method:"POST",
        body:JSON.stringify(data),
        headers:{
            "Content-Type": "application/json"
        }
    }).then(res=>res.json()).then(res=>{
        layer.msg(res.msg)
        if(callback instanceof Function)callback(res);
    })
}
function curlGet(url,data,callback){
    url += getUrl(data)
    return fetch(url,{
        method:"GET",
        // body:JSON.stringify(data),
        headers:{}
    }).then(res=>res.json()).then(res=>{
        layer.msg(res.msg)
        if(callback instanceof Function)callback(res);
    })
}

function getUrl(obj){
    let arr = Object.keys(obj)
    let result = '?'
    for(i=0;i<arr.length;i++){
        result += `${arr[i]}=${obj[arr[i]]}`
        if(i<arr.length-1){
            result += `&`
        }
    }
    return result
}