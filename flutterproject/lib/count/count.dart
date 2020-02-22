import 'package:flutter/material.dart';

class Count extends StatelessWidget {
  Count({Key key, this.count}) : super(key: key);
  final int count;

  //box样式
  final boxDecoration = BoxDecoration(
    color: Colors.white,
    borderRadius: BorderRadius.circular(190.0),
  );

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: BoxDecoration(color: Colors.grey),
      child: Container(
        padding: EdgeInsets.all(20),
        child: DecoratedBox(
          decoration: boxDecoration,
          child: Container(
            padding: EdgeInsets.all(20),
            margin: EdgeInsets.all(10),
            child: buildColumn(),
          ),
        ),
      ),
    );
  }

  Column buildColumn() {
    return Column(children: <Widget>[
      Text(
        '曾经有一份真挚的感情，我却没有珍惜，如果能再来一次我会说我爱你，如果非要给这份爱加上一个期限，我希望是 $count 万年',
        style: buildTextStyle(),
        textAlign: TextAlign.center,
        maxLines: 8,
      ),
    ]);
  }

  TextStyle buildTextStyle() {
    return TextStyle(
      fontSize: 18.0,
      height: 1.4,
      fontFamily: 'cao',
      color: Colors.blue,
    );
  }
}
