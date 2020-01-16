var api = [{
    url: "/admin/user/list", //用户列表
    config: {
        method: "GET",
        body: {
            page:1,
            limit:10
        },
        headers:'{"Content-Type": "application/json"}'
    }
},
{
    url: "/admin/user/addUser", //用户列表
    config: {
        method: "POST",
        body: {
            name:"sukeai",
            password:"sukeaiz"
        },
        headers:'{"Content-Type": "application/json"}'
    }
},
{
    url: "/admin/user/updateUser", //用户列表
    config: {
        method: "POST",
        body: {
            name:"sukdajsd",
            id:"119"
        },
        headers:'{"Content-Type": "application/json"}'
    }
},
{
    url: "/admin/user/delUser", //用户列表
    config: {
        method: "POST",
        body: {
            id:"119"
        },
        headers:'{"Content-Type": "application/json"}'
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
        headers:'{"Content-Type": "application/json"}'
    }
}]
