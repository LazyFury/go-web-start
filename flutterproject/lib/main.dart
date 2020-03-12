import 'package:flutter/material.dart';

import 'layout.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  MyApp({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: "APP",
      home: Home(),
      theme: ThemeData(primaryColor: Colors.blue),
      themeMode: ThemeMode.light,
    );
  }
}

class Home extends StatelessWidget {
  const Home({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Layout(
        title: "Detail",
        isTabbar: true,
        child: Column(
          children: <Widget>[
            Text('center'),
            FlatButton(
                onPressed: () {
                  print(Navigator.canPop(context));
                  Navigator.push(
                    context,
                    new MaterialPageRoute(
                      builder: (context) => Layout(
                        child: Text('data'),
                      ),
                    ),
                  );
                },
                child: Text("detail"))
          ],
        ));
  }
}
