import React from 'react'
import { Image } from 'react-native'

export function Icon(props) {
    return (
        <Image
            style={{ width: 30, height: 30 }}
            defaultSource={require("../../../assets/icon.png")}
            {...props}
        ></Image>
    )
}