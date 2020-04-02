import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutterproject/utils/color.dart';
import 'package:flutterproject/utils/utils.dart';

import 'App.dart';

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

class Global {
  static Future init() async {
    // WidgetsFlutterBinding.ensureInitialized();

    print("init");
    await Utils.setTimeout(Duration(milliseconds: 1000), () {
      print("loading success");
    });

    // int i = 0;
    // Utils.setInterval(Duration(milliseconds: 1000), (t) {
    //   if (i < 10) {
    //     print(i++);
    //   } else {
    //     t.cancel();
    //     t = null;
    //   }
    // });
  }
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
