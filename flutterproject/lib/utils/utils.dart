import 'dart:async';

class Utils {
  static Future setTimeout(Duration d, Function callback) {
    return new Future.delayed(d, () {
      callback();
    });
  }

  static setInterval(Duration d, void Function(Timer) callback) {
    return new Timer.periodic(d, callback);
  }
}
