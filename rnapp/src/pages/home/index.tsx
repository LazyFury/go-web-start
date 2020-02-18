import React from 'react';
import { SafeAreaView, StatusBar, Text, Button } from 'react-native';
import List from '../../components/list/list';
import Item from '../../components/list/defaultItem'
import Navbar from '../../components/custom/navbar';
// import Video from 'react-native-video'
export default function App({ navigation }) {
    function jump(Item = {}) {
        console.log('item组件回调方法获取到Item文本：', Item, this.navigationOptions)
        navigation.navigate("User")
    }

    return (
        <SafeAreaView style={{ flex: 1, backgroundColor: "#fff" }}>
            <StatusBar></StatusBar>
            {/* 导航条 */}
            <Navbar
                title="标题洒大大大大啊大大都是点点滴滴"
                borderLine={true}
            ></Navbar>
            {/* <Video source={{ uri: "" }}></Video> */}
            {/* 列表  */}
            <List name="列表"
                ListHeaderComponent={false}
                renderItem={(props) => Item({ ...props, callback: jump })}></List>
        </SafeAreaView>
    )
}

App.navigationOptions = {
    title: "hello World!",
    headerShown: false,
}