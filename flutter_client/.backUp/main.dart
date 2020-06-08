import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';
// import 'package:flutterproject/count/count.dart';
import 'package:flutterproject/item.dart';
import 'package:flutterproject/page/detail.dart';

import 'count/count.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      home: Home(title: "App 标题栏"),
      theme: ThemeData(
        primaryColor: Colors.yellow,
      ),
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
  int _tabbarcurrentIndex = 0;

  tabbarTap(e) {
    print(e);
    setState(() {
      _tabbarcurrentIndex = e;
    });
  }

  getPage(i) {
    return tabbarPages[i];
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

    var bottomNavigationBar2 = BottomNavigationBar(
      items: <BottomNavigationBarItem>[
        BottomNavigationBarItem(icon: Icon(Icons.history), title: Text('home')),
        BottomNavigationBarItem(icon: Icon(Icons.history), title: Text('info')),
        BottomNavigationBarItem(icon: Icon(Icons.history), title: Text('user')),
      ],
      currentIndex: _tabbarcurrentIndex,
      onTap: tabbarTap,
      // fixedColor: Colors.blue,
    );

    return new Scaffold(
      appBar: appBar2,
      drawer: drawer2,
      bottomNavigationBar: bottomNavigationBar2,
      body: getPage(_tabbarcurrentIndex),
      // body: Column(
      //   children: <Widget>[
      //     // Container(child: Count(count: count)),
      //     Expanded(
      //       child: listView(
      //           5,
      //           Container(
      //             child: buildSwiper(),
      //             height: 150,
      //           ),
      //           Text('footer')),
      //     )
      //   ],
      // )
    );
  }

  List<Widget> tabbarPages = new List.from([
    HomePage(),
    Row(
      children: <Widget>[Text('info1')],
    ),
    Row(
      children: <Widget>[Text('user')],
    ),
  ]);

  //  leading
  FlatButton leading() {
    return FlatButton(
      onPressed: () {},
      child: Icon(Icons.history, color: Colors.black),
    );
  }

  // actions
  List<Widget> actions(BuildContext context) {
    return <Widget>[
      FlatButton(
        onPressed: () {
          Navigator.push(
            context,
            new MaterialPageRoute(builder: (context) => Detail()),
          );
        },
        child: Text('Tabbar 按钮1'),
      ),
    ];
  }
}

class HomePage extends StatefulWidget {
  const HomePage({
    Key key,
  }) : super(key: key);

  @override
  HomePageSatus createState() => HomePageSatus();
}

class HomePageSatus extends State<HomePage> {
  int count = 0;

  updateCount() {
    setState(() {
      count++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Container(child: Count(count: count)),
        Expanded(
          child: listView(
              5,
              Container(
                child: buildSwiper(),
                height: 150,
              ),
              Text('footer')),
        )
      ],
    );
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
      layout: SwiperLayout.DEFAULT,
    );
  }

  ListView listView(int n, Widget header, Widget footer) {
    var list = new List<String>.generate(n, (i) => "列表子项标题第 $i 个");
    list.add('list_footer');
    list.insert(0, 'header');
    print(list);
    return ListView.builder(
        itemCount: n + 2,
        itemBuilder: (BuildContext context, int index) {
          if (index == 0 && list[0] == 'header') {
            return header;
          }
          if (list[index] == 'list_footer') {
            return footer;
          }
          return Item(
              click: () {
                updateCount();
                // Navigator.push(context,
                //     new MaterialPageRoute(builder: (context) => Detail()));
              },
              name: list[index]);
        });
  }
}
