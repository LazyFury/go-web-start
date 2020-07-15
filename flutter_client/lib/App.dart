import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';

import 'components/easyUse.dart';
import 'components/safeMode.dart';
import 'components/tabbar.dart';
import 'page/tabbar/cart.dart';
import 'page/tabbar/cate.dart';
import 'page/tabbar/home.dart';
import 'server/server.dart';

class App extends StatefulWidget {
  const App({
    Key key,
  }) : super(key: key);

  @override
  AppStatus createState() => AppStatus();
}

class AppStatus extends State<App> {
  int current = 0;

  void getInfo() {
    Http.get("/api/v1/posts").then((res) {
      print("请求成功");
      print(res);
    }).catchError((err) {
      print(err);
    });
  }

  initState() {
    super.initState();
    getInfo();
  }

  Widget pages() {
    List<Widget> pageList = [
      Home(),
      Cate(),
      Cart(),
      page('user'),
    ];
    return pageList[current];
  }

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.grey[100],
      child: Column(
        children: <Widget>[
          Expanded(
            child: pages(),
          ),
          Column(children: <Widget>[
            Tabbar(onChange: (i) {
              setState(() {
                current = i;
              });
            }),
            partation(height: 1),
            safeBottom(context, color: Colors.white)
          ])
        ],
      ),
    );
  }

  Widget page(String title) {
    var boxDecoration = BoxDecoration(
      color: Colors.white,
      borderRadius: BorderRadius.all(Radius.circular(10)),
      boxShadow: [
        BoxShadow(
          color: Colors.blue[100],
          offset: Offset(1, 3),
          blurRadius: 2,
        ),
        BoxShadow(
          color: Colors.blue[400],
          offset: Offset(-1, 3),
          blurRadius: 2,
        )
      ],
    );

    return AnnotatedRegion(
      value: SystemUiOverlayStyle.light,
      child: Column(
        children: <Widget>[
          Container(
            height: 200,
            width: double.infinity,
            child: Stack(
              children: [
                Positioned(
                  left: 0,
                  right: 0,
                  top: 0,
                  child: Column(
                    children: <Widget>[
                      safeStatusBar(context, color: Colors.blue),
                      Container(
                        decoration: BoxDecoration(color: Colors.blue),
                        height: 100,
                        width: double.infinity,
                        child: Text(""),
                      ),
                      safeStatusBar(context, color: Colors.transparent)
                    ],
                  ),
                ),
                Positioned(
                    bottom: 10,
                    left: 10,
                    right: 10,
                    child: Container(
                      decoration: boxDecoration,
                      height: 100,
                      width: double.infinity,
                      margin: EdgeInsets.all(20),
                      padding: EdgeInsets.all(10),
                      child: Text(
                        "hello world! there is a new flutter app demo positioned,let`s test it.cool! it`s runing!",
                        style: TextStyle(
                            fontSize: 20,
                            color: Colors.black,
                            fontFamily: "Regular"),
                        maxLines: 3,
                        overflow: TextOverflow.clip,
                      ),
                    )),
              ],
            ),
          ),
          Expanded(
            child: EasyRefresh(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  Text(title, style: TextStyle(fontSize: 80, fontFamily: "cao"))
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}
