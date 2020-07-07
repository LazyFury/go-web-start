import React, { useState } from 'react';
import { useLocation, history } from 'umi';
import { PageHeader } from 'antd';

export default ()=>{
    let param:any = useLocation()
    let {id} = param.query
    let [isEdit] = useState(Boolean(id))
   
    return <div>
        <PageHeader
        className="site-page-header fff"
        onBack={() => history.go(-1)}
        title={isEdit?"修改文章":"发布文章"}
        subTitle=""
      />
        add post
        {id}
    </div>
}