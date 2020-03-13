import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';

import 'App.dart';

void main() => runApp(MyApp());

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
