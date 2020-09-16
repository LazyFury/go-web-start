import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/widgets/easyUse.dart';
import 'package:flutterproject/widgets/layout.dart';
import 'package:flutterproject/widgets/navbar.dart';
import 'package:flutterproject/widgets/touchView.dart';
import 'package:flutterproject/library/SlidingEventsStatus.dart';
import 'package:flutterproject/utils/statusBar.dart';
import 'package:flutterproject/utils/utils.dart';

class Cart extends StatefulWidget {
  @override
  CartStatue createState() => CartStatue();
}

class CartStatue extends State<Cart> {
  bool isEdit = false;
  bool canClose = true;

  // 商品列表
  List<Map<String, dynamic>> cartListData = [
    {"name": "购物车商品3", "select": false},
    {"name": "购物车商品1", "select": false},
    {"name": "购物车商品2", "select": false}
  ];

// 是否以全选
  bool get selectAll => (() {
        var list = cartListData.where((e) => e['select']).toList();
        return list.length == cartListData.length;
      })();

  // 全选商品列表
  void selectAllCart() {
    setState(() {
      bool checked = selectAll;
      for (int i = 0; i < cartListData.length; i++) {
        cartListData[i]['select'] = !checked;
      }
    });
  }

  @override
  initState() {
    super.initState();
    eventBus.on<SlidingEventsBus>().listen((event) {
      switch (event.event) {
        case "inOperation":
          setState(() {
            canClose = false;
          });
          break;
        default:
      }
    });
  }

  // 重置侧滑部件
  resetAllSliding() {
    Utils.setTimeout(Duration(milliseconds: 50), () {
      if (canClose) {
        // print("全局重置 $canClose");
        eventBus.fire(new SlidingEventsBus("reset"));
      }
      setState(() {
        canClose = true;
      });
    });
  }

// 猜你喜欢
  List guessYouloveItData = ["asd"];

  @override
  Widget build(BuildContext context) {
    var textStyle = TextStyle(color: Colors.white, fontSize: 16);
    return GestureDetector(
      onPanDown: (e) {
        resetAllSliding();
      },
      child: Layout(
        navbar: buildNavbar(context, textStyle),
        child: Column(
          children: <Widget>[
            Expanded(
              child: EasyRefresh(
                onRefresh: () async {},
                child: Column(
                  children: <Widget>[
                    Column(
                      children: cartList(), //购物车列表
                    ),
                    Column(
                      children: <Widget>[
                        guessYouloveIt(),
                      ],
                    ), //猜你喜欢
                  ],
                ),
              ),
            ),
            settlement(), //结算
          ],
        ),
      ),
    );
  }

  // 猜你喜欢
  Widget guessYouloveIt() {
    if (guessYouloveItData.length == 0) {
      return Container(child: null);
    }
    return Text("猜你喜欢");
  }

  // 购物车列表
  List<Widget> cartList() {
    if (cartListData.length == 0) {
      return [noData(height: 300, title: "购物车暂无内容")];
    }
    return cartListData.asMap().entries.map((e) {
      return Column(
        children: <Widget>[cartItem(e), partation()],
      );
    }).toList();
  }

  Widget cartItem(MapEntry<int, Map<String, dynamic>> e) {
    return SlidingEvents(
      height: 120,
      leftChild: SlidingBackground(
        child: TouchView(
          child: Container(
              decoration: BoxDecoration(color: Colors.black54),
              child: Center(child: Text("collect"))),
        ),
        width: 100,
      ),
      rightChild: SlidingBackground(
        child: TouchView(
          child: Container(
            decoration: BoxDecoration(color: Colors.red),
            child: Center(
              child: Text("delete"),
            ),
          ),
        ),
        width: 120,
      ),
      child: Container(
        width: screenSize(context).width,
        height: 120,
        decoration: BoxDecoration(color: Colors.white),
        padding: EdgeInsets.symmetric(vertical: 10),
        child: Row(
          children: <Widget>[
            Checkbox(
                value: e.value['select'] == true,
                onChanged: (val) {
                  setState(() {
                    cartListData[e.key]['select'] = val;
                  });
                }),
            Container(
                width: 80,
                height: 80,
                child: networkImage(
                    "http://wx1.sinaimg.cn/mw600/9b61e9edgy1gdfzcposrrj20m80m50vm.jpg",
                    needLoading: true)),
            Expanded(
              child: Container(
                height: 80,
                padding: EdgeInsets.symmetric(horizontal: 10),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: <Widget>[
                    Text(
                      e.value['name'],
                      maxLines: 1,
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  // 结算
  Container settlement() => Container(
        decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.all(Radius.circular(999))),
        width: double.infinity,
        margin: EdgeInsets.all(10),
        padding: EdgeInsets.only(left: 20),
        child: Row(
          children: <Widget>[
            Expanded(
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: <Widget>[
                  Container(
                      child: !isEdit
                          ? Text(
                              "¥100320.09",
                              style: TextStyle(color: Colors.red, fontSize: 24),
                              maxLines: 1,
                            )
                          : null,
                      margin: EdgeInsets.only(right: 10)),
                  Text(
                    "共13件",
                    maxLines: 1,
                  )
                ],
              ),
            ),
            TouchView(
                onTap: () {},
                child: Container(
                  decoration: BoxDecoration(
                      color: Colors.red[500],
                      borderRadius:
                          BorderRadius.horizontal(right: Radius.circular(999))),
                  height: 60,
                  padding: EdgeInsets.fromLTRB(30, 0, 50, 0),
                  child: Center(
                      child: Text(isEdit ? "删除" : "结算",
                          style: TextStyle(color: Colors.white, fontSize: 18))),
                ))
          ],
        ),
      );

  // 自定义Navbar
  Widget buildNavbar(BuildContext context, TextStyle textStyle) {
    return navbar(context, title: "购物车", leftButton: [
      FlatButton(
          onPressed: selectAllCart,
          child: Text("${selectAll ? '全不选' : '全选'}", style: textStyle))
    ], rightButton: [
      FlatButton(
          onPressed: () {
            setState(() {
              isEdit = !isEdit;
            });
          },
          child: Text(isEdit ? "取消" : "编辑", style: textStyle))
    ]);
  }
}
