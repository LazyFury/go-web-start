import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/widgets/easyUse.dart';
import 'package:flutterproject/widgets/layout.dart';
import 'package:flutterproject/widgets/touchView.dart';
import 'package:flutterproject/page/product/list.dart';

import 'searchBar.dart';

class Cate extends StatefulWidget {
  @override
  CateStatus createState() => CateStatus();
}

class CateStatus extends State<Cate> {
  final cates = [
    {
      'name': "服饰1呀",
      'tmenu': [
        {
          'name': "夏装",
          'tmenu': [
            {'name': "上衣", 'tmenu': []},
            {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []}
          ]
        },
        {'name': "夏装", 'tmenu': []},
        {
          'name': "秋装",
          'tmenu': [
            {'name': "裤子", 'tmenu': []},
            // {'name': "裤子", 'tmenu': []},
            // {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []},
            {'name': "裤子", 'tmenu': []}
          ]
        },
        {'name': "冬装", 'tmenu': []}
      ]
    },
    {'name': "包包", 'tmenu': []},
    {
      'name': "日常",
      'tmenu': [
        {'name': "分类1", 'tmenu': []},
        {'name': "分类2", 'tmenu': []}
      ]
    }
  ];

  int current = 0;

  @override
  Widget build(BuildContext context) {
    return Layout(
      statusMode: SystemUiOverlayStyle.dark,
      title: "分类",
      navbar: Container(
        decoration: BoxDecoration(boxShadow: []),
        child: Column(
          children: <Widget>[searchBar(context), partation(height: 1)],
        ),
      ),
      child: Container(
        decoration: BoxDecoration(color: Colors.grey[200]),
        width: double.infinity,
        height: double.infinity,
        child: Row(
          children: <Widget>[
            firstClassify(),
            moreClassify(),
          ],
        ),
      ),
    );
  }

  // 一级分类
  Container firstClassify() => Container(
      width: 120,
      decoration: BoxDecoration(color: Colors.grey[100], boxShadow: []),
      child: EasyRefresh(
          child: Column(
              children: cates
                  .asMap()
                  .entries
                  .map((e) => TouchView(
                        onTap: () {
                          setState(() {
                            current = e.key;
                          });
                        },
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            Container(
                              width: double.infinity,
                              decoration: BoxDecoration(),
                              padding: EdgeInsets.symmetric(vertical: 16),
                              child: Row(
                                children: <Widget>[
                                  Container(
                                    width: 5,
                                    height: 30,
                                    margin: EdgeInsets.only(right: 10),
                                    decoration: BoxDecoration(
                                        color: e.key == current
                                            ? Colors.blue
                                            : Colors.grey[100]),
                                  ),
                                  Text(
                                    e.value['name'],
                                    style: TextStyle(
                                        fontSize: 18,
                                        color: e.key == current
                                            ? Colors.blue
                                            : Colors.black38),
                                    maxLines: 1,
                                    overflow: TextOverflow.clip,
                                  ),
                                ],
                              ),
                            ),
                            partation(height: 1, color: Colors.grey[200]),
                          ],
                        ),
                      ))
                  .toList())));
  // 更多分类
  Expanded moreClassify() => Expanded(
        child: EasyRefresh(
          child: Column(
            children: tmenu(),
          ),
        ),
      );
  // 二级分类
  List<Widget> tmenu() {
    var list = List<Map<String, Object>>.from(cates[current]['tmenu']);
    if (list.length == 0) {
      return [noData(height: 300, title: "暂无内容")];
    }
    return list.asMap().entries.map((e) {
      return Container(
        padding: EdgeInsets.fromLTRB(10, 0, 10, 20),
        width: double.infinity,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            Container(
              padding: EdgeInsets.symmetric(vertical: 10),
              child: Text(
                "" + e.value['name'],
                style: TextStyle(fontSize: 18),
              ),
            ),
            Wrap(
              children: thirdMenu(e.value['tmenu']),
              spacing: 20,
              runSpacing: 10,
              alignment: WrapAlignment.start,
              runAlignment: WrapAlignment.start,
              crossAxisAlignment: WrapCrossAlignment.start,
            )
          ],
        ),
      );
    }).toList();
  }

  // 三级分类
  List<Widget> thirdMenu(tmenu) {
    var list = List<Map<String, Object>>.from(tmenu);
    if (list.length == 0) {
      return [Container(child: Center(child: noData(title: "暂无下级分类")))];
    }
    return list.asMap().entries.map(
      (e) {
        return Container(
          margin: EdgeInsets.fromLTRB(0, 0, 0, 0),
          child: TouchView(
            onTap: () => toProductListPage(context),
            child: Column(
              children: <Widget>[
                Container(
                  decoration: BoxDecoration(
                      color: Colors.grey[300],
                      borderRadius: BorderRadius.all(Radius.circular(6))),
                  width: 40,
                  height: 40,
                ),
                Text(e.value['name']),
              ],
            ),
          ),
        );
      },
    ).toList();
  }

  // mehtods
  // 到商品列表页
  void toProductListPage(context) {
    Navigator.push(
        context, new MaterialPageRoute(builder: (context) => ProductList()));
  }
}
