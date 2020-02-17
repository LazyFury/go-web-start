/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow
 */
import 'react-native-gesture-handler'


import { createAppContainer } from 'react-navigation';
import { createStackNavigator } from 'react-navigation-stack';

import router from './src/router/router'


const AppNavigator = createStackNavigator({
    ...router
});

export default createAppContainer(AppNavigator);
