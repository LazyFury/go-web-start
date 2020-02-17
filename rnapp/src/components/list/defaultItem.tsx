import React from 'react'
import { TouchableHighlight, Image, View, Text } from "react-native"
import globalStyle from "../../static/style"
import { Icon } from "../icon"
import style from "./style"



/**
 * 列表默认 Flatlist randerItem
 * @param param0
 */
export default function DefaultItem({ item, index, callback = (item) => { } }) {
    const methods = {
        clickItem() {
            if (typeof callback === 'function') {
                callback(item)
            } else {
                console.info("回调函数错误")
            }
        }
    }
    return (
        <TouchableHighlight
            key={index}
            onPress={methods.clickItem}
            activeOpacity={.95}
        >
            <View style={[style.item, globalStyle.row]}>

                <Image style={{ width: 50, height: 50 }}
                    defaultSource={require("../../static/1.jpg")}
                    source={{ uri: "https://tva1.sinaimg.cn/large/0082zybpgy1gbulubh0scj31fy0tddle.jpg" }} ></Image>

                <View style={[globalStyle.col, { alignItems: "flex-start", justifyContent: "flex-start", marginLeft: 10, marginRight: 20 }]}>
                    <Text style={style.itemTitle}>Item标题</Text>
                    <Text style={{ color: "#999" }}>{item.test}</Text>
                </View>

                <Icon source={require("../../static/more.png")}></Icon>
            </View>
        </TouchableHighlight >
    )
}

