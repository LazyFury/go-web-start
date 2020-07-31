import 'package:flutter/material.dart';

IconData baseIconData(int code) {
  return IconData(code, fontFamily: "iconfont", matchTextDirection: false);
}

class IconFont {
  static IconData home = baseIconData(0xe6a9);
  static IconData user = baseIconData(0xe6f1);
  static IconData cart = baseIconData(0xe727);

  static IconData search = baseIconData(0xe70f);
  static IconData scanQR = baseIconData(0xe714);
  static IconData back = baseIconData(0xe6ef);

  static IconData shoping = baseIconData(0xe720);
  static IconData address = baseIconData(0xe721);
  static IconData bluetoothOff = baseIconData(0xe697);
}
