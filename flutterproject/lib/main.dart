import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'App.dart';
import 'global.dart';
import 'utils/color.dart';

void main() {
  Global.init().then((e) {
    print("加载成功");
    return runApp(MyApp());
  }).then((e) {
    print("初始化完成");
    // flutter 强制竖屏
    SystemChrome.setPreferredOrientations([DeviceOrientation.portraitUp]);
  });
}

class MyApp extends StatelessWidget {
  MyApp({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: "APP",
      home: App(),
      theme: ThemeData(primaryColor: CustomTheme.primaryColor),
    );
  }
}
