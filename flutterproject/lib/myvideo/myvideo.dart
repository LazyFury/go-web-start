import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:video_player/video_player.dart';

class VideoPlayerScreen extends StatefulWidget {
  VideoPlayerScreen({Key key}) : super(key: key);

  @override
  VideoPlayerState createState() => VideoPlayerState();
}

class VideoPlayerState extends State<VideoPlayerScreen> {
  VideoPlayerController _controller;

  // Future<void> _futureInit;

  @override
  void initState() {
    // TODO: implement initState
    _controller =
        VideoPlayerController.network("http://127.0.0.1:8080/static/hello.mp4");
    print(_controller);
    // _futureInit = _controller.initialize();

    super.initState();
  }

  @override
  void dispose() {
    // TODO: implement dispose
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: null,
      builder: (context, snapshot) {
        return AspectRatio(
          aspectRatio: _controller.value.aspectRatio,
          child: VideoPlayer(_controller),
        );
      },
    );
  }
}
