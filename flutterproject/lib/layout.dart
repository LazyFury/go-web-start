import 'package:flutter/material.dart';
import 'components/navbar.dart';

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
    return Material(
      color: Colors.grey[100],
      child: body(),
    );
  }

  // 主体内容区
  Column body() {
    return Column(
      children: <Widget>[
        navbar(context, title: widget.title),
        Expanded(
          child: widget.child,
        )
      ],
    );
  }
}
