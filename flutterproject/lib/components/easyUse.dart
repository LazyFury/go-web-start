import 'package:flutter/material.dart';

// 间隔
Container partation({double height = 10, Color color}) => Container(
      decoration: BoxDecoration(
        color: color != null ? color : Colors.grey[100],
      ),
      height: height,
    );

// 网络图片
Widget networkImage(String src) {
  return Image.network(
    src,
    fit: BoxFit.cover,
    loadingBuilder:
        (BuildContext context, Widget child, ImageChunkEvent loadingProgress) {
      // print(loadingProgress);
      if (loadingProgress == null) return child;
      // loadingProgress = ImageChunkEvent(
      //   cumulativeBytesLoaded: 10,
      //   expectedTotalBytes: 100,
      // );
      return Container(
        decoration: BoxDecoration(color: Colors.grey[100]),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              CircularProgressIndicator(
                backgroundColor: Colors.grey[200],
                value: loadingProgress.expectedTotalBytes != null
                    ? loadingProgress.cumulativeBytesLoaded /
                        loadingProgress.expectedTotalBytes
                    : null,
              ),
              Container(
                margin: EdgeInsets.only(top: 10),
                child: Text(loadingProgress.expectedTotalBytes != null
                    ? "加载中..." +
                        ((loadingProgress.cumulativeBytesLoaded /
                                        loadingProgress.expectedTotalBytes *
                                        100)
                                    .toInt() /
                                100)
                            .toString() +
                        "%"
                    : ""),
              )
            ],
          ),
        ),
      );
    },
  );
}
