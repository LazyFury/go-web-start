import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'navbar.dart';

class Layout extends StatefulWidget {
  final String title;
  final Widget child;
  final Widget navbar;
  final SystemUiOverlayStyle statusMode;
  const Layout(
      {Key key,
      this.title = '',
      @required this.child,
      this.navbar,
      this.statusMode})
      : super(key: key);

  @override
  LayoutState createState() => LayoutState();
}

class LayoutState extends State<Layout> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Material(
        color: Colors.grey[100],
        child: AnnotatedRegion(
            value: widget.statusMode != null
                ? widget.statusMode
                : SystemUiOverlayStyle.light,
            child: body()),
      ),
    );
  }

  // 主体内容区
  Column body() {
    return Column(
      children: <Widget>[
        makeNavbar(),
        Expanded(
          child: widget.child,
        )
      ],
    );
  }

  Widget makeNavbar() => widget.navbar != null
      ? widget.navbar
      : Container(child: navbar(context, title: widget.title));
}
