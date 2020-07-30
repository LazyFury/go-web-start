import 'package:flutter/material.dart';
import 'package:flutterproject/page/tabbar/user.dart';

import 'widgets/easyUse.dart';
import 'widgets/safeMode.dart';
import 'widgets/tabbar.dart';
import 'page/tabbar/cart.dart';
import 'page/tabbar/cate.dart';
import 'page/tabbar/home.dart';

class App extends StatefulWidget {
  const App({
    Key key,
  }) : super(key: key);

  @override
  AppStatus createState() => AppStatus();
}

class AppStatus extends State<App> {
  int current = 0;

  initState() {
    super.initState();
  }

  List<Widget> page = [
    Home(),
    Cate(),
    Cart(),
    UserCenter(),
  ];

  List<TabbarIcon> tabbars = [
    TabbarIcon(name: "首页", icon: Icons.home),
    TabbarIcon(name: "分类", icon: Icons.fast_forward),
    TabbarIcon(name: "购物车", icon: Icons.card_travel),
    TabbarIcon(name: "我的", icon: Icons.account_circle),
  ];

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.grey[100],
      child: Column(
        children: <Widget>[
          Expanded(
            child: page[current],
          ),
          Column(children: <Widget>[
            Tabbar(
                onChange: (i) => setState(() => current = i), tabbars: tabbars),
            partation(height: 1),
            safeBottom(context, color: Colors.white)
          ])
        ],
      ),
    );
  }
}
