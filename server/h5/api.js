let headers = '{"Content-Type": "application/json"}'
var api = [{
    url: "/admin/login", //用户列表
    config: {
        method: "GET",
        body: {
           username:"qwe",
           password:"123456"
        },
        headers
    }
},{
    url: "/admin/user/list", //用户列表
    config: {
        method: "GET",
        body: {
            page:1,
            limit:10,
        },
        headers
    }
},
{
    url: "/admin/user/addUser", //用户列表
    config: {
        method: "POST",
        body: {
            name:"sukeai",
            password:"sukeaiz",
        },
        headers
    }
},
{
    url: "/admin/user/updateUser", //用户列表
    config: {
        method: "POST",
        body: {
            name:"sukdajsd",
            id:"208",
            email:"2568597007@qq.com"
        },
        headers
    }
},
{
    url: "/admin/user/delUser", //用户列表
    config: {
        method: "POST",
        body: {
            id:"119"
        },
        headers
    }
}
,
{
    url: "/admin/user/repeatOfName", //用户列表
    config: {
        method: "GET",
        body: {
            name:"sukeai"
        },
        headers
    }
}
,
{
    url: "/admin/user/repeatOfEmail", //用户列表
    config: {
        method: "GET",
        body: {
            email:"2568597007@qq.com"
        },
        headers
    }
}
,
{
    url: "/wechat/jsApiConfig", //用户列表
    config: {
        method: "GET",
        body: {
            url:"http://abadboy.cn"
        },
        headers
    }
}
,
{
    url: "/wechat/wechat_redirect", //用户列表
    config: {
        method: "GET",
        body: {},
        headers
    }
}
,
{
    url: "/wechat/login", //用户列表
    config: {
        method: "GET",
        body: {
            code:"http://abadboy.cn"
        },
        headers
    }
},
{
    url: "/wechat/info", //用户列表
    config: {
        method: "GET",
        body: {
            id:"6"
        },
        headers
    }
}]
