import React, { useState } from 'react';
import { SafeAreaView, StatusBar, Text, View, Button } from 'react-native';
import List from '../../components/list/list';
import Item from '../../components/list/defaultItem'
import Navbar from '../../components/custom/navbar';
import MyVideo from '../../components/custom/video';
import globalStyle from '../../static/style';
import useTheme from '../../hooks/theme/useTheme';

export default function App({ navigation }) {
    function jump(Item = {}) {
        console.log('item组件回调方法获取到Item文本：', Item, this.navigationOptions)
        navigation.navigate("User", { item: Item })
    }
    let [dark, setDark] = useState(false)
    let [theme, setTheme] = useTheme()
    return (
        <SafeAreaView style={{ flex: 1, backgroundColor: "#fff" }}>
            <StatusBar backgroundColor={theme.background.level1} barStyle={dark ? "light-content" : "dark-content"}></StatusBar>
            <View style={[{ position: 'relative', left: 0, backgroundColor: theme.background.level1, width: "100%", height: 120, marginTop: -120 }]}></View>
            {/* 导航条 */}
            <Navbar
                title="标题洒大大大大啊大大都是点点滴滴"
                borderLine={false}
            ></Navbar>
            <Button title='切换主题' onPress={() => {
                setDark(!dark)
                setTheme(!dark ? "white" : "")
            }}></Button>
            <MyVideo style={{}}></MyVideo>
            {/* 列表  */}
            <List name="列表"
                ListHeaderComponent={false}
                renderItem={(props) => Item({ ...props, callback: jump, style: { backgroundColor: theme.Colors.Primary.level1 } })}></List>
        </SafeAreaView>
    )
}

App.navigationOptions = {
    title: "hello World!",
}
