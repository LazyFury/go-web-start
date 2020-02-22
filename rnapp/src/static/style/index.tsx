import React from 'react'
import { StyleSheet } from 'react-native'


let divFakeLine = {
    height: 1,
    backgroundColor: "#eee"
}

const globalStyle = StyleSheet.create({
    row: {
        display: 'flex',
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center"
    },
    col: {
        display: 'flex',
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center"
    },
    divFakeLine: {
        ...divFakeLine
    },
    divFakeLine10: {
        ...divFakeLine,
        height: 10
    }
})

export default globalStyle