import 'package:barcode_scan/barcode_scan.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutterproject/utils/color.dart';
import 'package:flutterproject/utils/iconFont.dart';
import 'package:flutterproject/widgets/safeMode.dart';
import 'package:flutterproject/widgets/touchView.dart';
import 'package:flutterproject/page/search.dart';
import 'package:google_fonts/google_fonts.dart';
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
            padding: EdgeInsets.fromLTRB(10, 0, 20, 0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: <Widget>[
                //选择地址
                addrView(),
                //输入框
                borderRadiusInput(context),
                // 扫码
                Container(
                  child: (scan ? scanView(context) : null),
                )
              ],
            ),
          ),
        ],
      ),
    );

TouchView addrView() {
  return TouchView(
    child: Container(
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.end,
        children: <Widget>[
          Icon(
            IconFont.address,
            color: Colors.green,
            size: 20,
          ),
          Text(
            '东莞',
            style: TextStyle(fontSize: 14),
          ),
        ],
      ),
    ),
  );
}

TouchView scanView(context) {
  return TouchView(
    onTap: () async {
      var result = await BarcodeScanner.scan(options: ScanOptions());
      print(result);
    },
    child: Icon(IconFont.scanQR, color: CustomTheme.primaryColor, size: 28),
  );
}

Expanded borderRadiusInput(context) {
  return Expanded(
    child: TouchView(
      onTap: () => toSearch(context),
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 10),
        margin: EdgeInsets.fromLTRB(10, 0, 10, 0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Icon(IconFont.search, size: 18, color: Colors.grey[600]),
            Padding(
              padding: const EdgeInsets.only(left: 4),
              child: Text(
                '搜索您想要找的商品名称～ ',
                style: GoogleFonts.peddana(
                    textStyle: TextStyle(color: Colors.grey, fontSize: 12)),
              ),
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

// @Methods
toSearch(BuildContext context) {
  Navigator.push(context, MaterialPageRoute(builder: (BuildContext context) {
    return SearchPage(context);
  }));
}
