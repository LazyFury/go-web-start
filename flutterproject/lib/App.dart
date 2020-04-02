import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutterproject/components/easyUse.dart';
import 'package:flutterproject/components/tabbar.dart';
import 'package:flutterproject/page/tabbar/home.dart';
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

  void getInfo() {
    // Http.get("admin_").then((res) {
    //   print("请求成功");
    //   print(res);
    // }).catchError((err) {
    //   print(err);
    // });
  }

  initState() {
    super.initState();
    getInfo();
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
      child: Column(children: <Widget>[
        Container(
            height: 200,
            decoration: BoxDecoration(color: CustomTheme.primaryColor),
            child: Stack(
                alignment: Alignment.center,
                fit: StackFit.expand,
                children: <Widget>[
                  Positioned(
                      child: networkImage(
                        "http://ww1.sinaimg.cn/mw600/a6fec82cgy1gct4jpt9u4j20wi1cqkjm.jpg",
                      ),
                      top: 0,
                      left: 0,
                      right: 0),
                  Center(
                    child: Column(
                      children: <Widget>[
                        safeStatusBar(context, color: Colors.transparent),
                        Expanded(
                          child: Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: <Widget>[
                              Container(
                                decoration: boxDecoration,
                                height: 100,
                                width: double.infinity,
                                margin: EdgeInsets.symmetric(horizontal: 20),
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
                              ),
                            ],
                          ),
                        )
                      ],
                    ),
                  )
                ])),
        Expanded(
          child: Column(children: <Widget>[
            Text(title, style: TextStyle(fontSize: 80, fontFamily: "Regular")),
          ]),
        )
      ]),
    );
  }
}
