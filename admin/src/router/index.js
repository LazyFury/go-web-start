import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


let pages = []
let requireComponents = require.context('../pages/',true,/\.vue$/)
requireComponents.keys().forEach(path=>{
	let name = /\/(.*?).vue/.exec(path)[1]
  let com = requireComponents(path)
  let component  = com.default || com
	pages.push({
    path:"/"+name,
    component
  })
})

console.log(pages)



const router = new Router({
  routes: [
    // 入口页面
    {
      path:"/",
      component:{
        template:"<div>1</div>",
        beforeCreate(){
          window.location.hash = "#/welcome"
        }
      }
    },
    // 业务页面
    ...pages,
    // 错误页面
    {
      path:"/404",
      component:{
        template:"<div>404</div>",
        created(){
          console.log('hello nofund')
        }
      }
    },
    // 为匹配到的页面会走到这里
    {
      path:"*",
      component:{
        template:"<div></div>",
        created(){
          console.log('hello nofund')
          window.location.hash = "#/404"
        }
      }
    }
  ]
})
export default router
