import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';
import 'package:flutterproject/utils/statusBar.dart';

class Layout extends StatefulWidget {
  const Layout({
    Key key,
    this.title = '',
    this.child,
  }) : super(key: key);
  final String title;
  final Widget child;
  @override
  LayoutState createState() => LayoutState();
}

class LayoutState extends State<Layout> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: body(),
    );
  }

  // 主体内容区
  Column body() {
    return Column(
      children: <Widget>[
        navbar(),
        Expanded(
          child: widget.child,
        )
      ],
    );
  }

  // 是否显示返回按钮
  Widget getBackButton() {
    if (ModalRoute.of(context).canPop) {
      return BackButton(color: CustomTheme.primaryTextColor);
    }
    return Row();
  }

  // navbar
  Widget navbar() {
    return Column(children: <Widget>[
      // 状态栏
      safeStatusBar(context),
      // navbar
      DecoratedBox(
        decoration: BoxDecoration(color: CustomTheme.primaryColor),
        child: Container(
          height: 50,
          width: double.infinity,
          child: Row(
            children: <Widget>[
              getBackButton(),
              Text(statusBarHeight(context).toString()),
              Text(Navigator.canPop(context).toString()),
            ],
          ),
        ),
      )
    ]);
  }
}

// 顶部状态安全位置
DecoratedBox safeStatusBar(context, {Color color}) => safeBox(
    height: statusBarHeight(context),
    color: color != null ? color : CustomTheme.primaryColor);
// 全面屏手机 底部安全位置
DecoratedBox safeBottom(context, {Color color}) => safeBox(
    height: bottomBarHeight(context),
    color: color != null ? color : CustomTheme.primaryColor);

DecoratedBox safeBox({double height, Color color = Colors.white}) {
  return DecoratedBox(
      decoration: BoxDecoration(color: color),
      child: Container(height: height));
}
