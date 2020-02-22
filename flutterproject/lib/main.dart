import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';
import 'package:flutterproject/count/count.dart';
import 'package:flutterproject/item.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      home: Home(title: "HOme TiTLE"),
      theme: ThemeData(primaryColor: Colors.green),
      title: "APP",
    );
  }
}

// Home state
class Home extends StatefulWidget {
  Home({Key key, this.title}) : super(key: key);

  final String title;

  @override
  HomeState createState() => HomeState();
}

// home 组件
class HomeState extends State<Home> {
  int count = 0;

  updateCount() {
    setState(() {
      count++;
    });
  }

  @override
  Widget build(BuildContext context) {
    var appBar2 = AppBar(
      title: Text(widget.title),
      centerTitle: true,
      leading: leading(),
      actions: actions(context),
    );

    var drawer2 = Row(
      children: <Widget>[Container(child: Text("row2"))],
    );

    return new Scaffold(
        appBar: appBar2,
        drawer: drawer2,
        body: Column(
          children: <Widget>[
            Container(
              child: buildSwiper(),
              height: 150,
            ),
            Container(child: Count(count: count)),
            Expanded(
              child: listView(5),
            )
          ],
        ));
  }

  Swiper buildSwiper() {
    return new Swiper(
        itemCount: 3,
        itemBuilder: (BuildContext context, int index) {
          return Image.network(
            "http://via.placeholder.com/350x150",
            fit: BoxFit.fill,
          );
        },
        pagination: new SwiperPagination(),
        control: new SwiperControl(),
        itemWidth: 300,
        layout: SwiperLayout.DEFAULT);
  }

  ListView listView(int n) {
    var list = new List<String>.generate(n, (i) => "imtendn $i");

    return ListView.builder(
        itemCount: n,
        itemBuilder: (BuildContext context, int index) {
          return Item(click: () {}, name: list[index]);
        });
  }

  //  leading
  FlatButton leading() {
    return FlatButton(
      onPressed: () {},
      child: Icon(Icons.history, color: Colors.white),
    );
  }

  // actions
  List<Widget> actions(BuildContext context) {
    return <Widget>[
      FlatButton(
        onPressed: () {},
        child: Text('data'),
      ),
    ];
  }
}
