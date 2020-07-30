import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';

import 'easyUse.dart';

Swiper buildSwiper() {
  return new Swiper(
    itemCount: 3,
    itemBuilder: (BuildContext context, int index) {
      return Container(
        decoration: BoxDecoration(color: Colors.grey[100]),
        height: 180,
        child: networkImage(
            "http://ww1.sinaimg.cn/mw600/a6fec82cgy1gct4jpt9u4j20wi1cqkjm_err.jpg",
            needLoading: true),
      );
    },
    pagination: new SwiperPagination(),
    // control: new SwiperControl(),
    // itemWidth: 300,
    layout: SwiperLayout.DEFAULT,
  );
}
