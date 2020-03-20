import 'dart:convert';
import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutterproject/components/easyUse.dart';
import 'package:flutterproject/components/tabbar.dart';
import 'package:flutterproject/page/tabbar/home.dart';
import 'package:flutterproject/server/Http.dart';
import 'package:flutterproject/utils/color.dart';

import 'components/safeMode.dart';

class App extends StatefulWidget {
  const App({
    Key key,
  }) : super(key: key);

  @override
  AppStatus createState() => AppStatus();
}

class AppStatus extends State<App> {
  int current = 0;

  void getInfo() async {
    var c = HttpClient();
    var uri = Uri(
      scheme: 'http',
      host: 'go.abadboy.cn',
      path: '/admin/login',
    );
    var req = await c.postUrl(uri);
    req.headers.add('token', 'value');
    req.add(utf8.encode({'name': "suke"}.toString()));
    var res = await req.close();
    print(await res.transform(Utf8Decoder()).join());
    print(uri.toString());
    Http.get().then((res) => {print(res)});
  }

  initState() {
    super.initState();
    // getInfo();
  }

  Widget pages() {
    List<Widget> pageList = [
      Home(),
      page('info'),
      page('hotel'),
      page('user'),
    ];
    return pageList[current];
  }

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.transparent,
      child: DecoratedBox(
        decoration: BoxDecoration(color: Colors.grey[100]),
        child: Column(
          children: <Widget>[
            Expanded(
              child: pages(),
            ),
            Column(
              children: <Widget>[
                Tabbar(
                  onChange: (i) {
                    setState(() {
                      current = i;
                    });
                  },
                ),
                safeBottom(context, color: Colors.white)
              ],
            )
          ],
        ),
      ),
    );
  }

  Widget page(String title) {
    return AnnotatedRegion(
      value: SystemUiOverlayStyle.light,
      child: Column(children: <Widget>[
        Container(
          height: 200,
          decoration: BoxDecoration(color: CustomTheme.primaryColor),
          child: Stack(
            alignment: Alignment.center,
            fit: StackFit.expand,
            children: <Widget>[
              networkImage(
                  "http://ww1.sinaimg.cn/mw600/a6fec82cgy1gct4jpt9u4j20wi1cqkjm.jpg"),
              Positioned(
                child: safeStatusBar(context, color: Colors.transparent),
                top: 0,
                left: 0,
                right: 0,
              )
            ],
          ),
        ),
        Expanded(
          child: Column(children: <Widget>[
            Text(title, style: TextStyle(fontSize: 80)),
          ]),
        )
      ]),
    );
  }
}
