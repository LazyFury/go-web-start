import React from 'react'
import { View, Text, Button, StatusBar, NativeModules } from 'react-native'
import { ScrollView } from 'react-native-gesture-handler'
import useTheme from '../../hooks/theme/useTheme'

let d = NativeModules.CalendarManager



export default function User({ route, navigation }) {
    let [theme] = useTheme()
    let { item } = route.params
    console.log(navigation)
    return (
        <ScrollView>
            <StatusBar barStyle="dark-content"></StatusBar>
            <Text style={{ color: theme.Colors.Primary.level1 }}>
                {JSON.stringify(item)}
                asds
            </Text>
            <Button title='native' onPress={() => {
                console.log(d)
                d.addEvent("suke", "asd")
            }}></Button>
            <Button title='back' onPress={() => {
                navigation.goBack()
            }}></Button>
        </ScrollView >
    )
}