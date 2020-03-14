// 顶部状态安全位置
import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';
import 'package:flutterproject/utils/statusBar.dart';

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
