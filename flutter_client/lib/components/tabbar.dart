import 'package:flutter/material.dart';
import '../components/touchView.dart';

class Tabbar extends StatefulWidget {
  const Tabbar({Key key, this.onChange}) : super(key: key);

  final onChange;
  @override
  TabbarState createState() => TabbarState();
}

class TabbarState extends State<Tabbar> {
  List tabbars = [
    {'name': '首页', 'icon': Icons.home},
    {'name': '分类', 'icon': Icons.receipt},
    {'name': '购物车', 'icon': Icons.card_travel},
    {'name': '我的', 'icon': Icons.account_circle},
  ];
  int current = 0;

  Row buildRow() {
    return Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: tabbars
            .asMap()
            .map((i, e) {
              return MapEntry(
                i,
                tabbarItem(
                  current == i,
                  i,
                  name: e['name'],
                  icon: e['icon'],
                ),
              );
            })
            .values
            .toList());
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 54,
      child: DecoratedBox(
        decoration: BoxDecoration(color: Colors.white),
        child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              buildRow(),
            ]),
      ),
    );
  }

  Widget tabbarItem(bool selected, int i, {icon, name}) {
    var color = selected ? Colors.blue : Colors.black;
    return TouchView(
      onTap: () {
        setState(() {
          current = i;
          widget.onChange(i);
        });
      },
      child: Column(
        children: <Widget>[
          Icon(
            icon,
            color: color,
            size: 24,
          ),
          Text(
            name,
            style: TextStyle(color: color, fontSize: 16),
          ),
        ],
      ),
    );
  }
}
