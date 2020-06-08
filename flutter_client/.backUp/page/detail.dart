import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';

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
  BodyBuilder({
    Key key,
  }) : super(key: key) {
    print('BodyBuilder Init');
  }

  @override
  Widget build(BuildContext context) {
    return EasyRefresh(
      onRefresh: () async {
        print('onRefresh');
      },
      child: Center(
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
      ),
    );
  }
}
