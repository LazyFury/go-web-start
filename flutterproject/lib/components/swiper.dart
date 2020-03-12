import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';

Swiper buildSwiper() {
  return new Swiper(
    itemCount: 3,
    itemBuilder: (BuildContext context, int index) {
      return Image.network(
        "http://via.placeholder.com/350x150",
        fit: BoxFit.fill,
      );
    },
    pagination: new SwiperPagination(),
    control: new SwiperControl(),
    itemWidth: 300,
    layout: SwiperLayout.DEFAULT,
  );
}
