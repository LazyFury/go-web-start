import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/widgets/easyUse.dart';
import 'package:flutterproject/widgets/layout.dart';
import 'package:flutterproject/widgets/swiper.dart';
import 'package:flutterproject/widgets/touchView.dart';
import 'package:flutterproject/page/product/detail.dart';
import 'package:flutterproject/page/tabbar/searchBar.dart';
import 'package:flutterproject/server/server.dart';
import 'package:flutterproject/utils/statusBar.dart';

class Home extends StatefulWidget {
  @override
  HomeState createState() => HomeState();
}

class HomeState extends State<Home> {
  @override
  initState() {
    super.initState();
    Http.get("/api/v1/posts");
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      title: "Home",
      navbar: searchBar(context, scan: true),
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
                  header(),
                  partation(height: 1),
                  Container(
                      decoration: BoxDecoration(color: Colors.white),
                      padding: EdgeInsets.symmetric(vertical: 10),
                      child: Center(
                        child: Wrap(
                          spacing: 10,
                          runSpacing: 10,
                          children:
                              List<Widget>.generate(10, (index) => item(index)),
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

  Widget header() {
    return Container(
      decoration: BoxDecoration(color: Colors.white),
      child: Column(
        children: <Widget>[
          Container(
            child: buildSwiper(),
            height: 180,
          ),
          menuView(),
          partation(height: 1),
          news(),
          partation(),
          ad(),
          partation(),
          listTitle()
        ],
      ),
    );
  }

  toProductDetail(int i) {
    Navigator.push(context, new MaterialPageRoute(builder: (context) {
      return ProductDetailView();
    }));
  }

  Widget item(int i) {
    return TouchView(
      onTap: () => toProductDetail(i),
      child: Container(
        width: (screenSize(context).width - 50) / 2,
        decoration: BoxDecoration(color: Colors.white),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            ClipRRect(
              borderRadius: BorderRadius.circular(10.0),
              child: Container(
                height: (screenSize(context).width - 50) / 2,
                width: (screenSize(context).width - 50) / 2,
                child: networkImage(
                    'http://wxt.sinaimg.cn/mw600/c37719ddgy1gh9ca7fqtjj20go0xc0w4.jpg',
                    needLoading: true),
              ),
            ),
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                Container(
                  margin: EdgeInsets.symmetric(vertical: 6),
                  child: Text(
                    '测试商品标题测试商品标题测试商品标题',
                    style: TextStyle(fontSize: 14),
                    maxLines: 2,
                  ),
                ),
                Text(
                  "\¥123.09",
                  style: TextStyle(
                      color: Colors.red,
                      fontSize: 20,
                      fontWeight: FontWeight.bold),
                )
              ],
            ),
            partation(height: 12, color: Colors.transparent)
          ],
        ),
      ),
    );
  }

  Container menuView() {
    return Container(
      padding: EdgeInsets.fromLTRB(10, 20, 10, 20),
      width: double.infinity,
      child: Wrap(
        spacing: 10,
        runSpacing: 10,
        alignment: WrapAlignment.start,
        runAlignment: WrapAlignment.start,
        crossAxisAlignment: WrapCrossAlignment.start,
        children: menuList(),
      ),
    );
  }

  Container news() {
    return Container(
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
      ),
    );
  }

  TouchView ad() {
    return TouchView(
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
          'http://wx1.sinaimg.cn/mw600/007Xv5XOgy1gcz8jttrsvj31400u0x6p.jpg',
        ),
      ),
    );
  }

  Container listTitle() {
    return Container(
      height: 50,
      padding: EdgeInsets.symmetric(horizontal: 20),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          Image.asset('static/image/farm.png', width: 30, height: 30),
          Expanded(
              child: Container(
            child: Text(
              '爆款推荐',
              style: TextStyle(fontSize: 16),
            ),
            margin: EdgeInsets.only(left: 10),
          ))
        ],
      ),
    );
  }

  List<Widget> menuList() {
    return [
      '复习鸡',
      "挖掘鸡",
      "战斗鸡",
      "蒸汽鸡",
      "攻击鸡",
      "航空母鸡",
      "蒸汽鸡",
    ]
        .map(
          (e) => menuIcon(
              image: networkImage(
                  "http://wx1.sinaimg.cn/mw600/0085KTY1gy1gh9gzuk9pkj30hs0vlwgu.jpg"),
              text: e),
        )
        .toList();
  }

  Widget menuIcon({@required Image image, @required String text}) => TouchView(
        onTap: () async {},
        child: Container(
          decoration: BoxDecoration(color: Colors.transparent),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Container(
                child: ClipRRect(
                  child: image,
                  borderRadius: BorderRadius.circular(10),
                ),
                height: 50,
                width: 50,
                margin: EdgeInsets.only(bottom: 2),
              ),
              Text(
                text,
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(fontSize: 12),
              )
            ],
          ),
          width: (screenSize(context).width - 60) / 5,
        ),
      );
}
