import 'utils/utils.dart';

class Global {
  static Future init() async {
    // WidgetsFlutterBinding.ensureInitialized();

    print("init");
    await Utils.setTimeout(Duration(milliseconds: 1000), () {
      print("loading success");
    });

    // int i = 0;
    // Utils.setInterval(Duration(milliseconds: 1000), (t) {
    //   if (i < 10) {
    //     print(i++);
    //   } else {
    //     t.cancel();
    //     t = null;
    //   }
    // });
  }
}
