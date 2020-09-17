import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/page/tabbar/user.dart';
import 'package:flutterproject/utils/iconFont.dart';

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

  @override
  initState() {
    super.initState();

    EasyRefresh.defaultHeader = ClassicalHeader(
        refreshReadyText: "松开刷新",
        refreshText: "下拉刷新",
        refreshedText: "刷新成功",
        refreshingText: "正在刷新",
        refreshFailedText: "请求失败",
        infoText: "最后更新于 %T");
  }

  List<Widget> pages = [
    Home(),
    Cate(),
    Cart(),
    UserCenter(),
  ];

  List<TabbarIcon> tabbars = [
    TabbarIcon(name: "首页", icon: IconFont.home),
    TabbarIcon(name: "商城", icon: IconFont.shoping),
    TabbarIcon(name: "购物车", icon: IconFont.cart),
    TabbarIcon(name: "我的", icon: IconFont.user),
  ];

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.grey[100],
      child: Column(
        children: <Widget>[
          Expanded(
            child: pages[current],
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
