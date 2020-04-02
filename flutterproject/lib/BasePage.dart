import 'package:flutter/material.dart';

class MyBasePage extends State<StatefulWidget> {
  final String title;

  MyBasePage({Key key, this.title}) {
    print("object");
  }

  @override
  Widget build(BuildContext context) {
    throw UnimplementedError("请重写build方法");
  }
}

class TestWidget extends MyBasePage {
  final String name;
  TestWidget({Key key, this.name}) {
    print("init class");
  }

  @override //初始化对象
  initState() {
    super.initState();
  }

  @override //当对象从当前树永久卸载
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      children: <Widget>[Text("hello")],
    );
  }
}
