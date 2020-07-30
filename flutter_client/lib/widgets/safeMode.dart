// 顶部状态安全位置
import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';
import 'package:flutterproject/utils/statusBar.dart';

Container safeStatusBar(context, {Color color}) => safeBox(
    height: safeArea(context).top,
    color: color != null ? color : CustomTheme.primaryColor);
// 全面屏手机 底部安全位置
Container safeBottom(context, {Color color}) => safeBox(
    height: safeArea(context).bottom,
    color: color != null ? color : CustomTheme.primaryColor);

Container safeBox({double height, Color color = Colors.white}) {
  return Container(
      decoration: BoxDecoration(color: color),
      child: Container(height: height));
}
