import 'package:event_bus/event_bus.dart';
import 'package:flutter/material.dart';
import '../components/touchView.dart';
import '../utils/utils.dart';

EventBus eventBus = new EventBus();

class SlidingEventsBus {
  final String event;

  SlidingEventsBus(this.event);
}

class SlidingEvents extends StatefulWidget {
  final Widget child;
  final double height;
  SlidingEvents({Key key, @required this.child, this.height});

  @override
  SlidingEventsStatus createState() => SlidingEventsStatus();
}

class SlidingEventsStatus extends State<SlidingEvents> {
  double leftOffset = 0;
  double rightOffset = 0;
  double target = 0;
  bool canClose = true;
  reset() {
    double length = leftOffset.abs();
    int step = 40;
    int milliseconds = length ~/ step;
    Utils.setInterval(
      Duration(milliseconds: milliseconds > 2 ? 2 : milliseconds),
      (t) {
        setState(() {
          leftOffset += (leftOffset > 0) ? -1 : 1;
          // print(leftOffset);
        });
        if (leftOffset.abs() - 1 <= 1) {
          leftOffset = 0;
          t.cancel();
          t = null;
        }
      },
    );
  }

  @override
  initState() {
    super.initState();
    eventBus.on<SlidingEventsBus>().listen((SlidingEventsBus event) {
      // print();
      if (event.event == "reset" && leftOffset != 0) {
        reset();
      }
    });
  }

  showConfrim({double end = 120}) {
    double length = (end - leftOffset.abs()).abs();
    int step = 80;
    int milliseconds = length ~/ step;
    Utils.setInterval(
      Duration(milliseconds: milliseconds > 2 ? 2 : milliseconds),
      (t) {
        setState(() {
          leftOffset += 1;
        });
        print(leftOffset);
        if (leftOffset < 0) {
          if (leftOffset.abs().floor() <= 120) {
            leftOffset = leftOffset > 0 ? 120 : -120;
            t.cancel();
            t = null;
          }
        } else {
          if (leftOffset.abs().floor() >= 120) {
            leftOffset = leftOffset > 0 ? 120 : -120;
            t.cancel();
            t = null;
          }
        }
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: widget.height,
      width: double.infinity,
      child: Stack(
        children: <Widget>[
          background(),
          Positioned(
            left: leftOffset,
            right: -leftOffset,
            child: GestureDetector(
              onPanUpdate: (e) {
                print(
                    "update:${e.localPosition.dx.toString()} target: $target ");
                var move = (e.localPosition.dx - target) * 0.75;
                setState(() {
                  print(leftOffset);
                  if (leftOffset.abs() >= 120) {
                    leftOffset = leftOffset > 0 ? 120 : -120;
                  } else {
                    leftOffset += move;
                  }
                  target = e.localPosition.dx;
                  print(leftOffset);
                });
              },
              onPanDown: (e) {
                if (leftOffset != 0) {
                  reset();
                }
              },
              onPanStart: (e) {
                Utils.setTimeout(Duration(milliseconds: 0), () {
                  // print(e);
                  setState(() {
                    target = e.localPosition.dx;
                  });
                });
              },
              onPanEnd: (e) {
                if (leftOffset.abs() > 100) {
                  showConfrim();
                } else {
                  reset();
                }
              },
              onPanCancel: () => reset(),
              child: widget.child,
            ),
          ),
        ],
      ),
    );
  }

  Positioned background() {
    return Positioned(
      child: TouchView(
        onTap: () {
          print("点击到了");
        },
        child: Container(
          height: widget.height,
          width: 120,
          decoration: BoxDecoration(color: Colors.red),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Text(
                "删除",
                style: TextStyle(color: Colors.white, fontSize: 18),
                maxLines: 1,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
