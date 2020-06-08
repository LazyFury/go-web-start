import 'package:flutter/material.dart';

ListView listView({int n, Widget item(info), Widget header, Widget footer}) {
  var list = new List<String>.generate(n, (i) => "列表子项标题第 $i 个");
  if (footer != null) {
    list.add('footer');
    n += 1;
  }
  if (header != null) {
    list.insert(0, 'header');
    n += 1;
  }
  // print(list);
  return ListView.builder(
      padding: EdgeInsets.all(0),
      itemCount: n,
      itemBuilder: (BuildContext context, int index) {
        if (index == 0 && list[0] == 'header') {
          return header;
        }
        if (list[index] == 'footer') {
          return footer;
        }
        return item(list[index]);
      });
}
