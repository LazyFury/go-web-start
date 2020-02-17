import React from 'react'
import { Image } from 'react-native'

export function Icon(props) {
    return (
        <Image
            style={{ width: 30, height: 30 }}
            defaultSource={require("../../static/1.jpg")}
            {...props}
        ></Image>
    )
}