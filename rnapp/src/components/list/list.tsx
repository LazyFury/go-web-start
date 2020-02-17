import React, { useState } from 'react'
import { userRequest } from '../../hooks/useRequest';
import { View, Text, FlatList, RefreshControl } from 'react-native';
import DefaultItem from './defaultItem';
import style from './style';

/**
 * 列表组件
 */
export default function List(props) {
    let { data, error, loading } = userRequest({
        api: "http://baidu.com"
    })
    let [state, setState] = useState(false)

    if (error) return <View><Text>something ornw</Text></View>
    // if (loading) return <View><Text>loading...</Text></View>
    return (
        <View style={style.list}>
            <FlatList
                ListHeaderComponent={() => ListHeader(props.name)}
                ListFooterComponent={() => loading != null && ListFooter(loading)}
                ListEmptyComponent={skeleton}
                data={data}
                keyExtractor={(item, index) => index.toString()}
                renderItem={DefaultItem}
                ItemSeparatorComponent={() => <View style={style.separator}></View>}
                refreshControl={
                    <RefreshControl
                        refreshing={state}
                        onRefresh={() => {
                            setState(true)
                            setTimeout(() => {
                                setState(false)
                            }, 1000);
                        }}
                    />
                }
                {...props}
            ></FlatList>
        </View>
    )
}

/**
 * 列表空
 */
function skeleton() {
    return <View style={style.item}><Text style={[style.HeaderTitle, { color: "#eee" }]}>loading...</Text></View>
}
/**
 * 页脚
 * @param loading 
 */
function ListFooter(loading: boolean) {
    return loading ? <Text>加载中</Text> : <Text>加载完成</Text>
}
/**
 * 页头
 * @param name 
 */
function ListHeader(name) {
    return (
        <View style={style.header}>
            <Text key='header' style={style.HeaderTitle}>{name || 'hello World!'}</Text>
        </View>

    )
}
