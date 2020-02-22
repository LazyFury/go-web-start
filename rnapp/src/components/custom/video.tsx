import React, { useState } from 'react'
import Video from 'react-native-video'
import globalStyle from '../../static/style';
import { View, StyleSheet, Button, Text } from 'react-native';

export default function MyVideo({ style }) {
    let [paused, setPaused] = useState(true) //默认暂停
    let player = null

    function videoError(e) {
        console.log(e)
        setPaused(true)
    }
    function onBuffer(e) {
        console.log(e)
    }

    return (
        <View style={[]}>
            <View style={[globalStyle.row, styles.videoBox, style]}>
                <Video source={{ uri: "http://127.0.0.1:8080/static/hello.m3u8" }}   // Can be a URL or a local file.
                    ref={(ref) => {
                        player = ref
                    }}                                      // Store reference
                    onBuffer={onBuffer}                // Callback when remote video is buffering
                    onError={videoError}               // Callback when video cannot be loaded
                    style={[{ flex: 1, width: 320, height: 240 }]}
                    paused={paused}
                    repeat={true}
                    resizeMode="cover"
                />
            </View>
            <View style={[globalStyle.row, styles.Controller]}>
                <View>
                    <Button color="#fff" onPress={() => setPaused(!paused)} title={paused ? "播放" : "暂停"}></Button>
                </View>
                <View>
                    <Button color='#fff' onPress={() => { }} title="全屏"></Button>
                </View>
            </View>
        </View >
    )
}

let styles = StyleSheet.create({
    videoBox: {
        height: 200, backgroundColor: "#333",
        overflow: "hidden"
    },
    Controller: {
        position: "absolute",
        width: "100%",
        left: 0,
        bottom: 0,
        backgroundColor: "#24292e7a",
        justifyContent: "space-between"
    },
    ControllerText: {
        color: "#fff"
    }
})