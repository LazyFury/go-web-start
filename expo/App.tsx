import React from 'react';
import { StyleSheet, Text, View, Button, StatusBar, SafeAreaView } from 'react-native';
import List from "./src/components/list/list"
import { NavigationContainer } from '@react-navigation/native'
import { createStackNavigator } from '@react-navigation/stack'
function Home() {
  return (
    <SafeAreaView style={{ backgroundColor: "#eee", flex: 1 }}>
      <StatusBar></StatusBar>
      <View>
        {/* <Button title='返回' onPress={() => { }} ></Button> */}
        <Text>navigion</Text>
      </View>
      <List name='列表标题'></List>
      <Text style={styles.tabbar}>Tabbar</Text>
    </SafeAreaView >
  )
}
let Stack = createStackNavigator()
export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen name="Home" component={Home} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  tabbar: {
    height: 80
  }
});
