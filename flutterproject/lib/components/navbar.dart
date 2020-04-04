// navbar
import 'package:flutter/material.dart';
import '../components/safeMode.dart';
import '../utils/color.dart';

Widget navbar(context,
    {String title: "标题", List<Widget> leftButton, List<Widget> rightButton}) {
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
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            slideBox(context,
                children:
                    leftButton != null ? leftButton : [getBackButton(context)]),
            // Text(statusBarHeight(context).toString()),
            Expanded(
              child: Padding(
                padding: EdgeInsets.all(10),
                child: Container(
                  child: Text(
                    title,
                    textAlign: TextAlign.center,
                    style: TextStyle(
                        color: Colors.white,
                        fontSize: 20,
                        fontWeight: FontWeight.bold),
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                    softWrap: true,
                  ),
                ),
              ),
            ),
            slideBox(context,
                children: rightButton != null ? rightButton : [],
                isleft: false),
          ],
        ),
      ),
    )
  ]);
}

Widget slideBox(context, {List<Widget> children, bool isleft: true}) =>
    Container(
      child: Row(
        mainAxisAlignment:
            isleft ? MainAxisAlignment.start : MainAxisAlignment.end,
        children: children,
      ),
      width: 100,
    );

// 是否显示返回按钮
Widget getBackButton(context) {
  if (ModalRoute.of(context).canPop) {
    return BackButton(color: CustomTheme.primaryTextColor);
  }
  return Row();
}
