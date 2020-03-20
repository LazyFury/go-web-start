import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/components/easyUse.dart';
import 'package:flutterproject/components/layout.dart';
import 'package:flutterproject/components/list.dart';
import 'package:flutterproject/components/safeMode.dart';
import 'package:flutterproject/components/swiper.dart';
import 'package:flutterproject/components/touchView.dart';
import 'package:flutterproject/page/tabbar/HomeComponents.dart';
import 'package:webview_flutter/webview_flutter.dart';

class Home extends StatefulWidget {
  @override
  HomeState createState() => HomeState();
}

class HomeState extends State<Home> {
  @override
  initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      title: "Home",
      navbar: searchBar(context),
      statusMode: SystemUiOverlayStyle.dark,
      child: Column(
        children: <Widget>[
          Expanded(
            child: EasyRefresh(
              onRefresh: () async {},
              onLoad: () async {
                print('onload');
              },
              child: Column(
                children: <Widget>[
                  Container(child: header()),
                  partation(height: 1),
                  Container(
                      decoration: BoxDecoration(color: Colors.white),
                      padding: EdgeInsets.symmetric(vertical: 10),
                      child: Center(
                        child: Wrap(
                          spacing: 10,
                          runSpacing: 10,
                          children:
                              List<Widget>.generate(10, (index) => item()),
                        ),
                      ))
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget item() {
    return TouchView(
      onTap: () {
        Navigator.push(context, new MaterialPageRoute(builder: (context) {
          return Layout(
              title: "商品详情",
              child: WebView(
                initialUrl: "http://baidu.com", //JS执行模式 是否允许JS执行
                javascriptMode: JavascriptMode.unrestricted,
              ));
        }));
      },
      child: Container(
        width: (screenSize(context).width - 50) / 2,
        decoration: BoxDecoration(color: Colors.white),
        child: Column(
          children: <Widget>[
            Container(
              height: (screenSize(context).width - 50) / 2,
              child: networkImage(
                  'https://lanhu.oss-cn-beijing.aliyuncs.com/SketchPngd21e5584e9c3b6891bb3df0be5ec52c4735705db31f63d4ecf566db18e7150b2'),
            ),
            Text('data'),
          ],
        ),
      ),
    );
  }

  Widget header() {
    return Container(
      decoration: BoxDecoration(color: Colors.white),
      child: Column(
        children: <Widget>[
          Container(
            child: buildSwiper(),
            height: 180,
          ),
          Container(
            padding: EdgeInsets.symmetric(vertical: 10),
            child: Center(
              child: Wrap(
                spacing: 10,
                runSpacing: 10,
                alignment: WrapAlignment.start,
                runAlignment: WrapAlignment.start,
                crossAxisAlignment: WrapCrossAlignment.start,
                children: <Widget>[
                  menuIcon(
                      image: AssetImage('static/image/camera.png'), text: '监控'),
                  menuIcon(
                      image: AssetImage('static/image/scan-fill.png'),
                      text: '扫码'),
                  menuIcon(
                      image: AssetImage('static/image/chicken.png'),
                      text: '复习鸡'),
                  menuIcon(
                      image: AssetImage('static/image/scan-fill.png'),
                      text: '监控'),
                  menuIcon(
                      image: AssetImage('static/image/chicken.png'),
                      text: '复习🐔'),
                  menuIcon(
                      image: AssetImage('static/image/camera.png'), text: '监控'),
                  menuIcon(
                      image: AssetImage('static/image/chicken.png'),
                      text: '监控'),
                ],
              ),
            ),
          ),
          partation(height: 1),
          Container(
              height: 40,
              padding: EdgeInsets.symmetric(horizontal: 20),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  Image.asset('static/image/news-icon.png', height: 20),
                  Expanded(
                      child: Container(
                    child: Text(
                      "瑶山富硒鸡拥有超高的营养价值，心动不如行动瑶山富硒鸡拥有超高的营养价值，心动不如行动",
                      style: TextStyle(color: Colors.grey),
                      maxLines: 1,
                      overflow: TextOverflow.ellipsis,
                    ),
                    margin: EdgeInsets.only(left: 10),
                  ))
                ],
              )),
          partation(),
          TouchView(
            onTap: () {
              Navigator.push(context, new MaterialPageRoute(builder: (context) {
                return Layout(
                    title: "广告详情",
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: <Widget>[Text('data')],
                    ));
              }));
            },
            child: Container(
              height: 100,
              width: double.infinity,
              child: networkImage(
                  'http://wx1.sinaimg.cn/mw600/007Xv5XOgy1gcz8jttrsvj31400u0x6p.jpg'),
            ),
          ),
          partation(),
          Container(
            height: 50,
            padding: EdgeInsets.symmetric(horizontal: 20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: <Widget>[
                Image.asset('static/image/farm.png', width: 30, height: 30),
                Expanded(
                    child: Container(
                  child: Text('爆款推荐'),
                  margin: EdgeInsets.only(left: 10),
                ))
              ],
            ),
          )
        ],
      ),
    );
  }

  Widget menuIcon({@required AssetImage image, @required String text}) =>
      TouchView(
        onTap: () async {},
        child: Container(
          decoration: BoxDecoration(color: Colors.transparent),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Container(
                child: Image(image: image, width: 50, height: 50),
                height: 50,
                width: 50,
              ),
              Text(text,
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: TextStyle(fontSize: 12))
            ],
          ),
          width: (screenSize(context).width - 80) / 5,
        ),
      );
}
