import React, { useState } from 'react'
import { userRequest } from '../../hooks/useRequest';
import { View, Text, FlatList, TouchableHighlight, StyleSheet, RefreshControl, Image } from 'react-native';

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
    if (loading === null) return <Text>上拉加载数据</Text>
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
/**
 * 列表默认 Flatlist randerItem
 * @param param0
 */
function DefaultItem({ item, index, separators }) {
    const methods = {
        clickItem() {
            console.log(item)
        }
    }
    return (
        <TouchableHighlight
            key={index}
            onPress={methods.clickItem}
            activeOpacity={.95}
        >
            <View style={style.item}>
                <Text style={style.itemTitle}>Item</Text>
                <Text>{item.test}</Text>
                <Image style={{ width: 50, height: 50 }} source={{ uri: "https://tva1.sinaimg.cn/large/0082zybpgy1gbulubh0scj31fy0tddle.jpg" }} ></Image>
            </View>
        </TouchableHighlight >
    )
}

const style = StyleSheet.create({
    list: {
        backgroundColor: "#ddd",
        flex: 1
    },
    header: {
        backgroundColor: "#ddd",
        flex: 1,
        alignItems: "flex-start",
        justifyContent: "center",
        height: 50,
        paddingHorizontal: 20,
        marginTop: 20
    },
    HeaderTitle: {
        fontSize: 32,
        fontWeight: "bold"
    },
    item: {
        backgroundColor: "#ffffff",
        flex: 1,
        justifyContent: "center",
        paddingHorizontal: 20,
        minHeight: 60,
        paddingVertical: 10
    },
    itemTitle: {
        fontSize: 18
    },
    separator: {
        height: 10,
        // backgroundColor: "#ddd"
    }
})