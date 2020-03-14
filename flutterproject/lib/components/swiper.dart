import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';

Swiper buildSwiper() {
  return new Swiper(
    itemCount: 3,
    itemBuilder: (BuildContext context, int index) {
      return Image.network(
        "http://ww1.sinaimg.cn/mw600/a6fec82cgy1gct4jpt9u4j20wi1cqkjm.jpg",
        fit: BoxFit.cover,
      );
    },
    pagination: new SwiperPagination(),
    control: new SwiperControl(),
    // itemWidth: 300,
    layout: SwiperLayout.DEFAULT,
  );
}
