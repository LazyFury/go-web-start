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
            "http://wx4.sinaimg.cn/mw600/a746f3dcly1gh9f8t536tj20u0140n6d.jpg",
            needLoading: true),
      );
    },
    pagination: new SwiperPagination(),
    // control: new SwiperControl(),
    // itemWidth: 300,
    layout: SwiperLayout.DEFAULT,
  );
}
