import React from 'react'
import { StyleSheet } from 'react-native'

let style = StyleSheet.create({
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

    separator: {
        height: 10,
        // backgroundColor: "#ddd"
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
})

export default style