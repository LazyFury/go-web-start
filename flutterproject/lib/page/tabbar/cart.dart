import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/components/easyUse.dart';
import 'package:flutterproject/components/layout.dart';
import 'package:flutterproject/components/navbar.dart';
import 'package:flutterproject/components/safeMode.dart';
import 'package:flutterproject/components/touchView.dart';
import 'package:flutterproject/library/SlidingEventsStatus.dart';
import 'package:flutterproject/utils/utils.dart';

class Cart extends StatefulWidget {
  @override
  CartStatue createState() => CartStatue();
}

class CartStatue extends State<Cart> {
  bool isEdit = false;
  bool canClose = true;

  List<Map<String, dynamic>> cartListData = [
    {"name": "购物车商品", "select": false},
    {"name": "购物车商品", "select": false},
    {"name": "购物车商品", "select": false}
  ];

  void selectAllCart() {
    setState(() {
      for (int i = 0; i < cartListData.length; i++) {
        cartListData[i]['select'] = true;
      }
    });
  }

  @override
  initState() {
    super.initState();
    // onPanDown 从父组件传到子组件，所以这里相应到要晚一点，重置方法需要加一点延时
    eventBus.on<SlidingEventsBus>().listen((event) {
      if (event.event == "inOperation") {
        print("操作按钮");
        setState(() {
          canClose = false;
        });
      }
    });
  }

  // 重置侧滑部件
  resetAllSliding() {
    Utils.setTimeout(Duration(milliseconds: 50), () {
      if (canClose) {
        print("全局重置$canClose");
        eventBus.fire(new SlidingEventsBus("reset"));
      }
      setState(() {
        canClose = true;
      });
    });
  }

  List guessYouloveItData = [];
  @override
  Widget build(BuildContext context) {
    var textStyle = TextStyle(color: Colors.white, fontSize: 18);
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
      leftChild:
          SlidingBackground(child: Center(child: Text("collect")), width: 0),
      rightChild:
          SlidingBackground(child: Center(child: Text("delete")), width: 220),
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
                    "http://wx1.sinaimg.cn/mw600/9b61e9edgy1gdfzcposrrj20m80m50vm.jpg")),
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
      FlatButton(onPressed: selectAllCart, child: Text("全选", style: textStyle))
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
