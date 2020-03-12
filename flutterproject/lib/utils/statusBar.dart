import 'package:flutter/material.dart';

// 获取状态栏高度
double statusBarHeight(BuildContext c) => MediaQuery.of(c).padding.top;
// 底部安全区域
double bottomBarHeight(BuildContext c) => MediaQuery.of(c).padding.bottom;
