import 'dart:convert';
import 'dart:io';
import 'dart:typed_data';

import 'package:flutter/foundation.dart';

Uri baseURIConfig({
  String scheme = "http",
  String userInfo,
  String host = "go.abadboy.cn",
  int port,
  String path = "",
  Iterable<String> pathSegments,
  String query,
  Map<String, dynamic> queryParameters,
  String fragment,
}) =>
    Uri(
      scheme: scheme,
      host: host,
      path: path,
      port: port,
      userInfo: userInfo,
      pathSegments: pathSegments,
      query: query,
      queryParameters: queryParameters,
      fragment: fragment,
    );

class Http {
  ///request
  static Future<dynamic> request(
    String method,
    Uri uri, {
    bool isFile = false,
  }) async {
    print(uri);
    return new HttpClient().openUrl(method, uri).then((req) async {
      // print(req);
      var res = await req.close().catchError((err) {
        print(err);
      });
      if (res.statusCode != HttpStatus.ok) {
        return Future.error("Fail 网络请求错误: ${req.uri}");
      }
      return _result(res, req, isFile);
    }).catchError((err) {
      throw "Err:" + err.toString();
    });
  }

  // Get请求 /
  static Future get(String path,
      {String query, Map<String, dynamic> data, bool isFile = false}) {
    return request(
        'get',
        baseURIConfig(
          path: path,
          query: query,
          queryParameters: data,
        ),
        isFile: isFile);
  }

  // Post请求
  static Future post(String path,
      {String query, Map<String, dynamic> data, bool isFile = false}) {
    return request(
        'post',
        baseURIConfig(
          path: path,
          query: query,
          queryParameters: data,
        ),
        isFile: isFile);
  }

  // 处理返回的结果
  static Future<dynamic> _result(
      HttpClientResponse res, HttpClientRequest req, bool isFile) async {
    // var res = await req.close();

    if (isFile) {
      return res;
    } else {
      String body;

      Uint8List data;
      try {
        data = await consolidateHttpClientResponseBytes(res);
      } catch (err) {
        print("转换bytes失败：" + err.toString());
      }

      try {
        // 文本类型返回
        body = utf8.decode(data);

        print("""\n
>>>>>>>>>>>>>>>>>httpClient.Success>>>>>>>>>>
Response.statusCode: [${res.statusCode.toString()}] 
Request.Uri: ${req.uri} 
Response.body: $body 
>>>>>>>>>>>>>>>>>""");

        if (body == null) {
          return Future.error("返回内容空");
        }

        dynamic bodyJson;
        // 尝试json解码
        try {
          bodyJson = json.decode(body);
        } catch (err) {
          return Future.error("json 解码失败");
        }
        return bodyJson;
      } catch (err) {
        print("转换文本失败：" + err.toString());
      }
    }
  }
}
