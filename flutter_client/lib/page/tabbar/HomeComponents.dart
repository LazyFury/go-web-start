import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutterproject/components/safeMode.dart';
import 'package:flutterproject/components/touchView.dart';
import 'package:flutterproject/page/search.dart';
// import 'utils/color.dart';

Widget searchBar(context, {bool scan = false}) => Container(
      decoration: BoxDecoration(
        color: Colors.white,
      ),
      child: Column(
        children: <Widget>[
          safeStatusBar(context, color: Colors.transparent),
          Container(
            height: 50,
            padding: EdgeInsets.fromLTRB(20, 0, 20, 0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Row(
                  children: <Widget>[
                    Image(
                        image: AssetImage('static/image/map.png'),
                        width: 20,
                        height: 20),
                    Text('东莞'),
                  ],
                ),
                borderRadiusInput(context),
                Container(
                  child: (scan
                      ? TouchView(
                          onTap: () {
                            Navigator.push(context, MaterialPageRoute(builder: (BuildContext context) {
                              return SearchPage(context);
                            }));
                          },
                          child: Image(
                              image: AssetImage('static/image/scan.png'),
                              width: 20,
                              height: 20),
                        )
                      : null),
                )
              ],
            ),
          ),
        ],
      ),
    );

Expanded borderRadiusInput(context) {
  return Expanded(
    child: TouchView(
      onTap: () {
        Navigator.push(context, MaterialPageRoute(builder: (BuildContext context) {
          return SearchPage(context);
        }));
      },
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 10),
        margin: EdgeInsets.fromLTRB(20, 0, 16, 0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Image(
              image: AssetImage('static/image/search.png'),
              width: 16,
              height: 16,
            ),
            Text(
              '搜索您想要找的商品名称',
              style: TextStyle(color: Colors.grey),
            ),
          ],
        ),
        height: 30,
        decoration: BoxDecoration(
          color: Colors.grey[200],
          borderRadius: BorderRadius.circular(20.0),
        ),
      ),
    ),
  );
}
