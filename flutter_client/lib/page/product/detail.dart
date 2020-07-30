import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/widgets/easyUse.dart';
import 'package:flutterproject/widgets/layout.dart';

class ProductDetailView extends StatefulWidget {
  @override
  _ProductDetailView createState() => _ProductDetailView();
}

class _ProductDetailView extends State<ProductDetailView> {
  @override
  Widget build(BuildContext context) {
    return Layout(
      title: "商品详情",
      child: EasyRefresh(
        onRefresh: () async {},
        child: Column(
          children: <Widget>[
            Container(
              child: networkImage(
                  "http://wx3.sinaimg.cn/mw600/44f2ef1bgy1gdg8mfzsfij21400u0qag.jpg"),
            ),
          ],
        ),
      ),
    );
  }
}
