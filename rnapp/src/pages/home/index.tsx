import React from 'react';
import { SafeAreaView, StatusBar, Text, Button } from 'react-native';
import List from '../../components/list/list';
import Item from '../../components/list/defaultItem'
import Navbar from '../../components/custom/navbar';

export default function App({ navigation }) {

    function jump(Item = {}) {
        console.log('item组件回调方法获取到Item文本：', Item)
        navigation.navigate("User")
    }

    return (
        <SafeAreaView style={{ flex: 1, backgroundColor: "#eee" }}>
            <StatusBar></StatusBar>
            {/* 导航条 */}
            <Navbar
                renderTitle={() => <Button onPress={() => { }} title='hasdsdsello'></Button>}
            ></Navbar>

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