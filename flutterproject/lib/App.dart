import 'dart:convert';
import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/components/list.dart';
import 'package:flutterproject/components/swiper.dart';
import 'package:flutterproject/components/tabbar.dart';
import 'package:flutterproject/components/touchView.dart';
import 'package:flutterproject/server/Http.dart';

import 'components/safeMode.dart';
import 'layout.dart';

class Home extends StatefulWidget {
  const Home({
    Key key,
  }) : super(key: key);

  @override
  HomeStatus createState() => HomeStatus();
}

class HomeStatus extends State<Home> {
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
    getInfo();
  }

  Widget pages() {
    List<Widget> pageList = [
      homePage(),
      page('info'),
      page('hotel'),
      page('user'),
    ];
    return pageList[current];
  }

  @override
  Widget build(BuildContext context) {
    return AnnotatedRegion(
      value: SystemUiOverlayStyle.light,
      child: Material(
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
      ),
    );
  }

  Widget page(String title) {
    return AnnotatedRegion(
      value: SystemUiOverlayStyle.dark,
      child: Center(
        child: Column(
          children: <Widget>[
            safeStatusBar(context, color: Colors.white),
            Expanded(
              child: Row(
                children: <Widget>[Text(title, style: TextStyle(fontSize: 30))],
                mainAxisAlignment: MainAxisAlignment.center,
              ),
            ),
          ],
        ),
      ),
    );
  }

// homePage
  Layout homePage() {
    return Layout(
      title: "首页",
      child: Column(
        children: <Widget>[
          // Text('data'),
          Expanded(
            child: EasyRefresh(
              onRefresh: () async {
                print('onrefresh');
              },
              // onLoad: () async {},
              child: listView(
                  n: 6,
                  item: (info) => buildItem(info),
                  header: Container(child: buildSwiper(), height: 150),
                  footer: Text('footer')),
            ),
          ),
        ],
      ),
    );
  }

  Widget buildItem(info) => TouchView(
        onTap: () {
          print('object');
        },
        child: Column(
          children: <Widget>[
            Container(
              padding: EdgeInsets.all(10),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: <Widget>[
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: <Widget>[
                        Text(
                          "hello world",
                          style: TextStyle(fontSize: 20),
                        ),
                        Text(info.toString()),
                      ],
                    ),
                  ),
                  Ink(
                    decoration: new BoxDecoration(
                      color: Colors.blue,
                      borderRadius: BorderRadius.all(new Radius.circular(99)),
                    ),
                    child: InkWell(
                      onTap: () {
                        print('inkwell');
                        Navigator.push(
                            context,
                            new MaterialPageRoute(
                                builder: (BuildContext context) => Layout(
                                      title: "详情页详情页详情页详情页详情页详情页详情页",
                                      child: Column(
                                        children: <Widget>[Text("detail")],
                                      ),
                                    )));
                      },
                      borderRadius: new BorderRadius.circular(25.0),
                      child: Container(
                        width: 80,
                        height: 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            Text(
                              'data',
                              style: TextStyle(color: Colors.white),
                            )
                          ],
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Container(
              width: double.infinity,
              height: 1,
              child: DecoratedBox(
                  decoration: BoxDecoration(color: Colors.grey[200])),
            )
          ],
        ),
      );
}
