/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow
 */
import React from 'react'
import 'react-native-gesture-handler'

import { NavigationContainer } from '@react-navigation/native';

import Router from './src/router/router'

function App() {
    return (
        <NavigationContainer>
            <Router></Router>
        </NavigationContainer>
    )
}

export default App;
