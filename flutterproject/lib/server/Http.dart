import 'dart:io';

class Http {
  static request({type: 'get'}) async {
    var client = new HttpClient();
    var req;
    var res;

    if (type == 'get') {
      req = await client.getUrl(Uri.parse('http://go.abadboy.cn'));
    }

    res = await req.close();
    return '请求结果:' + res.headers.toString();
  }

  static Future get() => request(
        type: 'get',
      );

  static Future post() {
    return request(type: 'post');
  }
}
