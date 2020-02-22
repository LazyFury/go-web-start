import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Item extends StatelessWidget {
  Item({Key key, this.click, this.name = 'hello world'}) : super(key: key);
  final Function click;
  final String name;

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: EdgeInsets.all(10),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: children(),
        ));
  }

  List<Widget> children() {
    return <Widget>[
      Text('data'),
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: <Widget>[Text(name), Text("row2")],
      ),
      FlatButton(
          onPressed: () {
            print("hello world button!");
            click();
          },
          child: Text('data'))
    ];
  }
}
