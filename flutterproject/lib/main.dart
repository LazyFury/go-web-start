import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';

import 'App.dart';

void main() {
  Global.init().then((e) => runApp(MyApp()));
}

class Global {
  static Future init() async {}
}

class MyApp extends StatelessWidget {
  MyApp({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: "APP",
      home: Home(),
      theme: ThemeData(primaryColor: CustomTheme.primaryColor),
    );
  }
}
