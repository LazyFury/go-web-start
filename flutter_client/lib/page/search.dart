import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:flutterproject/components/layout.dart';

class SearchPage extends StatelessWidget{
  SearchPage(BuildContext context);

  @override
  Widget build(BuildContext context) {
    return Layout(
      title:"hello world",
      child: EasyRefresh(onRefresh: ()async{},child: Text("hello this is a search page!"),),
    );
  }



}