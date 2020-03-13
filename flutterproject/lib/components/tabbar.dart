import 'package:flutter/material.dart';
import 'package:flutterproject/components/touchView.dart';

class Tabbar extends StatefulWidget {
  const Tabbar({Key key, this.onChange}) : super(key: key);

  final onChange;
  @override
  TabbarState createState() => TabbarState();
}

class TabbarState extends State<Tabbar> {
  List tabbars = [
    {'name': 'home', 'icon': Icons.home},
    {'name': 'info', 'icon': Icons.info},
    {'name': 'hotal', 'icon': Icons.hotel},
    {'name': 'user', 'icon': Icons.account_circle},
  ];
  int current = 0;

  Row buildRow() {
    List<Widget> list = [];

    for (var i = 0; i < tabbars.length; i++) {
      var item = tabbars[i];
      list.add(
        tabbarItem(
          current == i,
          i,
          name: item['name'],
          icon: item['icon'],
        ),
      );
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: list,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 54,
      child: DecoratedBox(
        decoration: BoxDecoration(color: Colors.grey[200]),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            buildRow(),
          ],
        ),
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
          Icon(icon, color: color),
          Text(
            name,
            style: TextStyle(color: color),
          ),
        ],
      ),
    );
  }
}
