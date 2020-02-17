import React, { useState } from 'react';
import { StyleSheet, Text, View, Button, StatusBar, SafeAreaView } from 'react-native';
import List from "./src/components/list/list"

export default function App() {
  return (
    <SafeAreaView style={{ backgroundColor: "#eee", flex: 1 }
    }>
      <StatusBar></StatusBar>
      <View>
        {/* <Button title='返回' onPress={() => { }} ></Button> */}
        <Text>navigion</Text>
      </View>
      <List name='列表标题'></List>
      <Text style={styles.tabbar}>Tabbar</Text>
    </SafeAreaView >
  );
}

const styles = StyleSheet.create({
  tabbar: {
    height: 80
  }
});
