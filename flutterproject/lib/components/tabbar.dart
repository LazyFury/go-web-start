import 'package:flutter/material.dart';

class Tabbar extends StatefulWidget {
  @override
  TabbarState createState() => TabbarState();
}

class TabbarState extends State<Tabbar> {
  @override
  Widget build(BuildContext context) {
    return Container(
      height: 54,
      child: DecoratedBox(
        decoration: BoxDecoration(color: Colors.grey[200]),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: <Widget>[Text('tabbar'), Text('tabbar')],
        ),
      ),
    );
  }
}
