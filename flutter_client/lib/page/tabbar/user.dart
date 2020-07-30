import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/widgets/safeMode.dart';
import 'package:lottie/lottie.dart';

class UserCenter extends StatefulWidget {
  @override
  _UserCenter createState() => _UserCenter();
}

class _UserCenter extends State<UserCenter> {
  @override
  Widget build(BuildContext context) {
    return page("asd");
  }

  Widget page(String title) {
    var boxDecoration = BoxDecoration(
      color: Colors.white,
      borderRadius: BorderRadius.all(Radius.circular(10)),
      boxShadow: [
        BoxShadow(
          color: Colors.blue[100],
          offset: Offset(1, 3),
          blurRadius: 2,
        ),
        BoxShadow(
          color: Colors.blue[400],
          offset: Offset(-1, 3),
          blurRadius: 2,
        )
      ],
    );

    return AnnotatedRegion(
      value: SystemUiOverlayStyle.light,
      child: Column(
        children: <Widget>[
          Container(
            height: 200,
            width: double.infinity,
            child: Stack(
              children: [
                Positioned(
                  left: 0,
                  right: 0,
                  top: 0,
                  child: Column(
                    children: <Widget>[
                      safeStatusBar(context, color: Colors.blue),
                      Container(
                        decoration: BoxDecoration(color: Colors.blue),
                        height: 100,
                        width: double.infinity,
                        child: Text(""),
                      ),
                      safeStatusBar(context, color: Colors.transparent)
                    ],
                  ),
                ),
                Positioned(
                    bottom: 10,
                    left: 10,
                    right: 10,
                    child: Container(
                      decoration: boxDecoration,
                      height: 100,
                      width: double.infinity,
                      margin: EdgeInsets.all(20),
                      padding: EdgeInsets.all(10),
                      child: Text(
                        "hello world! there is a new flutter app demo positioned,let`s test it.cool! it`s runing!",
                        style: TextStyle(
                            fontSize: 20,
                            color: Colors.black,
                            fontFamily: "Regular"),
                        maxLines: 3,
                        overflow: TextOverflow.clip,
                      ),
                    )),
              ],
            ),
          ),
          Expanded(
            child: EasyRefresh(
              onRefresh: () async {},
              child: Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  Text(title,
                      style: TextStyle(fontSize: 80, fontFamily: "cao")),
                  Lottie.asset("static/lottie/A.json")
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}
