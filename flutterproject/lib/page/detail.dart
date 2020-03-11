import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Detail extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("详情")),
      body: BodyBuilder(),
    );
  }
}

// 主体
class BodyBuilder extends StatelessWidget {
  const BodyBuilder({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Row(
        children: <Widget>[
          Text('data'),
          FlatButton(
              onPressed: () {
                Navigator.push(
                  context,
                  new MaterialPageRoute(builder: (context) => Detail()),
                );
              },
              child: Text("Detail"))
        ],
      ),
    );
  }
}
