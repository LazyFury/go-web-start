import 'package:flutter/material.dart';

// 获取状态栏高度
EdgeInsets safeArea(BuildContext c) => MediaQuery.of(c).padding;
// 屏幕尺寸
Size screenSize(context) => MediaQuery.of(context).size;
