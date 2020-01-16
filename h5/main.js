let baseUrl = "/api"
// Vue.component('VueJsonPretty',VueJsonPretty)
var app = new Vue({
    el: '#app',
    data() {
        return {
            message: 'Hello Vue!',
            code: 1,
            urlList: api,
            current:0,
            result: '{}',
            time:""
        }
    },
    methods: {
        send({url, config}) {
            // 拷贝 否则删除属性的时候会影响 data 
            let conf = JSON.parse(JSON.stringify(config))
            // Get方式请求
            let isGet = conf.method == "get"||conf.method == "GET"
            if (isGet){
                // 拼接url参数 删除 body 
                url += getUrl(conf.body)
                delete conf.body
            }
            //转换body
            if(conf.body){
                conf.body = JSON.stringify(conf.body)
            }
            // 转换headers 
            if (conf.headers) {
                conf.headers = JSON.parse(conf.headers)
            }
            console.log(conf)

            // 返回fetch请求 
            return fetch(baseUrl + url, {...conf}).then(res => {
                console.log(res)
                // 默认所有结果都是json类型
                return res.json()
            }).then(res => {
                console.log(res)
                this.message = res.msg
                this.code = res.code
                this.result = highLight(JSON.stringify(res, undefined, 3))
                this.time = new Date().toString()
            })
        }
    }
})


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

function highLight(json){
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
     return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
            var cls = 'number';
            if (/^"/.test(match)) {
                if (/:$/.test(match)) {
                    cls = 'key';
                } else {
                    cls = 'string';
                }
            } else if (/true|false/.test(match)) {
                cls = 'boolean';
            } else if (/null/.test(match)) {
                cls = 'null';
            }
            return '<span class="' + cls + '">' + match + '</span>';
        });
}

let list = {
    title:"基础服务",
    list:[
        {
            name:"标题",
            value:"阿斯顿哈看电话卡等哈看得见啊等哈说"
        },
        {
            name:"标题",
            value:"阿斯顿哈看电话卡等哈看得见啊等哈说"
        },
        {
            name:"标题",
            value:"阿斯顿哈看电话卡等哈看得见啊等哈说"
        }
    ]
}