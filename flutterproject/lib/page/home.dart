import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/components/list.dart';
import 'package:flutterproject/components/swiper.dart';
import 'package:flutterproject/components/touchView.dart';

import '../layout.dart';

class Home extends StatefulWidget {
  const Home({
    Key key,
  }) : super(key: key);

  @override
  HomeStatus createState() => HomeStatus();
}

class HomeStatus extends State<Home> {
  @override
  Widget build(BuildContext context) {
    return Layout(
      isTabbar: true,
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
