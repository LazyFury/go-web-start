import 'package:flutter/material.dart';
import 'package:flutterproject/utils/statusBar.dart';

class Layout extends StatefulWidget {
  const Layout({Key key, this.title = '', this.child, this.isTabbar = false})
      : super(key: key);
  final String title;
  final Widget child;
  final bool isTabbar;
  @override
  LayoutState createState() => LayoutState();
}

class LayoutState extends State<Layout> {
  tabbarWidget() {
    if (widget.isTabbar) {
      return Tabbar(context: context);
    }
    return Row();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: body(),
    );
  }

  Column body() {
    return Column(
      children: <Widget>[
        Navbar(context: context),
        Expanded(
          child: widget.child,
        ),
        tabbarWidget()
      ],
    );
  }
}

class Tabbar extends StatelessWidget {
  const Tabbar({
    Key key,
    @required this.context,
  }) : super(key: key);

  final BuildContext context;

  @override
  Widget build(BuildContext context) {
    return Column(children: <Widget>[
      DecoratedBox(
          decoration: BoxDecoration(color: Colors.white),
          child: Container(
            height: 50,
          )),
      // 状态栏
      DecoratedBox(
        decoration: BoxDecoration(color: Colors.white10),
        child: Container(
          height: bottomBarHeight(context),
          width: double.infinity,
          child: Row(
            children: <Widget>[Text(bottomBarHeight(context).toString())],
          ),
        ),
      ),
    ]);
  }
}

class Navbar extends StatelessWidget {
  const Navbar({
    Key key,
    @required this.context,
  }) : super(key: key);

  final BuildContext context;

  @override
  Widget build(BuildContext context) {
    return Column(children: <Widget>[
      // 状态栏
      DecoratedBox(
        decoration: BoxDecoration(color: Colors.deepOrange[600]),
        child: Container(
          height: statusBarHeight(context),
        ),
      ),
      // navbar
      DecoratedBox(
        decoration: BoxDecoration(color: Colors.deepOrange[500]),
        child: Container(
          height: 50,
          width: double.infinity,
          child: Row(
            children: <Widget>[
              Text(statusBarHeight(context).toString()),
              Text(Navigator.canPop(context).toString())
            ],
          ),
        ),
      )
    ]);
  }
}
