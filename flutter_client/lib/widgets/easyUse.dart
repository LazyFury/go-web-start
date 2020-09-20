import 'package:flutter/material.dart';
import '../library/NetImage.dart';

// 间隔
Container partation({double height = 10, Color color}) => Container(
      decoration: BoxDecoration(
        color: color != null ? color : Colors.grey[100],
      ),
      height: height,
    );
// 空数据组件
Widget noData({double height = 100, String title = ""}) => Container(
      height: height,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          Expanded(
            child: Container(
              height: height > 280 ? height / 2 : height - 30,
              child: Image(
                image: AssetImage(
                  "static/image/noData.png",
                ),
                fit: BoxFit.cover,
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(top: 10),
            child: Text(title),
          )
        ],
      ),
    );

/// 网络图片,设置了默认loadingBuilder和errBuilder
/// ```dart
///   Container(
///     width:100,
///     height:100,
///     child:networkImage("http://xxx.com/xx.jpg")
///   )
/// ```
Image networkImage(String src, {bool needLoading = false}) {
  ImageProvider image = new NetImage(src, asset: "static/image/empty.gif");
  return Image(
    image: image,
    fit: BoxFit.cover,
    repeat: ImageRepeat.noRepeat,
    loadingBuilder: needLoading ? loadingBuilder : null,
    errorBuilder: errBuilder,
  );
}

Widget errBuilder(BuildContext context, Object obj, StackTrace trace) {
  // print("no Image");
  return Image(
    image: AssetImage("static/image/empty.gif"),
    fit: BoxFit.cover,
    repeat: ImageRepeat.repeat,
  );
}

Widget loadingBuilder(
  BuildContext context,
  Widget child,
  ImageChunkEvent loadingProgress,
) {
  // print(loadingProgress);
  if (loadingProgress == null) return child;
  // loadingProgress = ImageChunkEvent(
  //   cumulativeBytesLoaded: 10,
  //   expectedTotalBytes: 100,
  // );
  return Container(decoration: BoxDecoration(color: Colors.grey[100]));

  // return Container(
  //   decoration: BoxDecoration(color: Colors.grey[100]),
  //   child: Center(
  //     child: Column(
  //       mainAxisAlignment: MainAxisAlignment.center,
  //       children: <Widget>[
  //         CircularProgressIndicator(
  //           backgroundColor: Colors.grey[200],
  //           value: loadingProgress.expectedTotalBytes != null
  //               ? loadingProgress.cumulativeBytesLoaded /
  //                   loadingProgress.expectedTotalBytes
  //               : null,
  //         ),
  //         Container(
  //           margin: EdgeInsets.only(top: 10),
  //           child: Text(loadingProgress.expectedTotalBytes != null
  //               ? "加载中...\n" +
  //                   ((loadingProgress.cumulativeBytesLoaded /
  //                                   loadingProgress.expectedTotalBytes *
  //                                   100)
  //                               .toInt() /
  //                           100)
  //                       .toString() +
  //                   "%"
  //               : ""),
  //         )
  //       ],
  //     ),
  //   ),
  // );
}
