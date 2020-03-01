import React from 'react'
import { View, Text, StyleSheet, Button } from 'react-native'
import globalStyle from '../../static/style'

export default function Navbar(props) {
    console.log(props)
    return (
        <View>
            <View style={[globalStyle.row, style.navbar]} {...props}>
                {props.renderLeft && <props.renderLeft /> || <Button onPress={() => { }} title="Left"></Button>}
                <View>
                    <Text style={style.title}>{props.title || "navbar title"}</Text>
                </View>
                {props.renderRight && <props.renderRight /> || <Button onPress={() => { }} title="Right"></Button>}
            </View>

            {/* 底部线条  */}
            {props.borderLine && <View style={globalStyle.divFakeLine10}></View>}
        </View>
    )
}

let style = StyleSheet.create({
    navbar: {
        height: 60,
        justifyContent: "space-between",
        backgroundColor: "#fff",
        paddingHorizontal: 20
    },
    title: {
        fontSize: 20,
        width: 180,
        overflow: "hidden",
        lineHeight: 30,
        height: 30
    }
})