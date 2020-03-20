import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';

import 'easyUse.dart';

Swiper buildSwiper() {
  return new Swiper(
    itemCount: 3,
    itemBuilder: (BuildContext context, int index) {
      return Container(
        decoration: BoxDecoration(color: Colors.grey[100]),
        child: networkImage(
            "http://wx3.sinaimg.cn/mw600/b5259065gy1gcz507dabkj20hs0hsgn3.jpg?date=202003191210123"),
      );
    },
    pagination: new SwiperPagination(),
    // control: new SwiperControl(),
    // itemWidth: 300,
    layout: SwiperLayout.DEFAULT,
  );
}
