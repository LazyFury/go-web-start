import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/utils/statusBar.dart';
import 'package:flutterproject/widgets/easyUse.dart';
import 'package:flutterproject/widgets/layout.dart';
import 'package:flutterproject/widgets/touchView.dart';

class ProductList extends StatefulWidget {
  @override
  ProductListState createState() => ProductListState();
}

class ProductListState extends State<ProductList> {
  List<int> products = [1, 2, 3, 4, 1, 23, 3];

  Future loadMore() async {
    setState(() {
      products.addAll(List<int>.generate(100, (index) => index + 1));
    });
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      title: "商品列表",
      child: EasyRefresh(
          onRefresh: () async {},
          onLoad: loadMore,
          child: Column(
            children: [
              Container(
                  child: networkImage("http://baidu.com"),
                  width: double.infinity,
                  height: 160),
              // 列表
              Container(
                child: Column(
                  children: List<Widget>.generate(
                    products.length,
                    (index) => item(products[index]),
                  ),
                ),
              )
            ],
          )),
    );
  }

  Widget item(int item) {
    return TouchView(
      onTap: () async {
        showXConfirm().then((sure) {
          print(sure);
        });
      },
      child: Container(
        decoration: BoxDecoration(color: Colors.white),
        padding: EdgeInsets.all(10),
        margin: EdgeInsets.fromLTRB(10, 10, 10, 5),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
                child: networkImage("http://baidu.com"),
                width: 100,
                height: 100),
            Expanded(
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Container(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        "Title !!",
                        style: TextStyle(fontSize: 18),
                      ),
                      Container(
                        child: Text(
                          "这里是简介内容的阿哈水库打火机快点哈科技的哈萨克就回答说",
                          style: TextStyle(color: Colors.grey),
                          overflow: TextOverflow.clip,
                          maxLines: 3,
                        ),
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.end,
                        children: [
                          Container(
                              height: 30,
                              child: IconButton(
                                icon: Icon(Icons.lock_outline),
                                onPressed: () {},
                                color: Colors.blue,
                              )),
                        ],
                      )
                    ],
                  ),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }

  Future<bool> showXConfirm() {
    Widget button(text, onTap) {
      return Expanded(
        child: TouchView(
          child: Container(
            height: 30,
            child: Center(
              child: Text(
                text,
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
            ),
          ),
          onTap: onTap,
        ),
      );
    }

    return showDialog<bool>(
      context: context,
      builder: (BuildContext context) {
        return Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Material(
                  color: Colors.transparent,
                  child: Container(
                    width: screenSize(context).width * .64,
                    decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(10)),
                    child: Column(
                      children: [
                        Text('title'),
                        Text("asdsdddd"),
                        Row(
                          children: [
                            button('text', () {}),
                            button('text', () {
                              Navigator.of(context).pop(true);
                            }),
                          ],
                        )
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ],
        );
      },
    );
  }
}
