import React from 'react'
import { View, Text, StyleSheet } from 'react-native'

export default function Navbar(props) {
    console.log(props)
    return (
        <View style={style.navbar} {...props}>
            {props.renderTitle && <props.renderTitle /> || <Text>Title</Text>}
        </View>
    )
}

let style = StyleSheet.create({
    navbar: {
        height: 60
    }
})