/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow
 */
import React from 'react'
import { createStackNavigator } from '@react-navigation/stack';
import App from '../pages/home'
import User from '../pages/home/user'

let Stack = createStackNavigator()
function Router() {
  return (
    <Stack.Navigator>
      <Stack.Screen name='Home' component={App} options={{ headerShown: false, }}></Stack.Screen>
      <Stack.Screen name='User' component={User}></Stack.Screen>
    </Stack.Navigator>
  )
}


export default Router
