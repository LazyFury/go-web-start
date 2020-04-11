import 'package:event_bus/event_bus.dart';
import 'package:flutter/material.dart';
import '../utils/utils.dart';

EventBus eventBus = new EventBus();

class SlidingEventsBus {
  final String event;

  SlidingEventsBus(this.event);
}

class SlidingBackground {
  final Widget child;
  final double width;
  SlidingBackground({Key key, this.child, this.width})
      : assert(width != null, child != null);
}

class SlidingEvents extends StatefulWidget {
  final Widget child;
  final double height;
  final SlidingBackground leftChild;
  final SlidingBackground rightChild;
  SlidingEvents(
      {Key key,
      @required this.child,
      this.height,
      this.leftChild,
      this.rightChild});

  @override
  SlidingEventsStatus createState() => SlidingEventsStatus();
}

class SlidingEventsStatus extends State<SlidingEvents> {
  double offset = 0;
  double target = 0;

  double get leftTarget => widget.leftChild.width;
  double get rightTarget => widget.rightChild.width;
  double get leftOffset => offset;
  double get rightOffset => -offset;
  bool get leftToRight => offset >= 0;
  bool get rightToLeft => offset <= 0;

  @override
  initState() {
    super.initState();
    onEvent();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: widget.height,
      width: double.infinity,
      child: Stack(
        children: <Widget>[
          // left
          Positioned(
            top: 0,
            left: 0,
            child:
                widget.leftChild != null ? background(widget.leftChild) : Row(),
          ),
          // right
          Positioned(
            top: 0,
            right: 0,
            child: widget.rightChild != null
                ? background(widget.rightChild)
                : Row(),
          ),
          // child
          Positioned(
            left: leftOffset,
            right: rightOffset,
            child: GestureDetector(
              behavior: HitTestBehavior.opaque,
              onPanUpdate: onPanUpdate,
              onPanDown: onPanDown,
              onPanStart: onPanStart,
              onPanEnd: onPanEnd,
              // onPanCancel: () => reset(),
              child: widget.child,
            ),
          ),
        ],
      ),
    );
  }

  void onEvent() {
    eventBus.on<SlidingEventsBus>().listen((SlidingEventsBus event) async {
      if (event.event == "reset" && offset != 0) {
        reset();
      }
    });
  }

  reset() {
    double length = offset.abs();
    int milliseconds = 20;
    int step = (length / milliseconds).floor().toInt();

    Utils.setInterval(
      Duration(milliseconds: milliseconds),
      (t) {
        setState(() {
          offset += (offset > 0) ? -step : step;
          // print(offset);
        });
        if (offset.abs() - step <= step) {
          offset = 0;
          t.cancel();
          t = null;
        }
      },
    );
  }

  showConfrim() {
    double length = 0;
    if (leftToRight) {
      length = leftTarget - offset;
    } else {
      length = rightTarget - offset;
    }
    int milliseconds = 20;
    int step = (length / milliseconds).floor().toInt();

    Utils.setInterval(
      Duration(milliseconds: milliseconds),
      (t) {
        setState(() {
          offset += offset > 0 ? step : -step;
        });
        print("$offset > $leftTarget || $offset < $rightTarget");
        if (offset > leftTarget || offset < -rightTarget && offset != 0) {
          offset = offset > 0 ? leftTarget : -rightTarget;
          t.cancel();
          t = null;
        }
      },
    );
  }

  void onPanEnd(e) {
    if (offset > leftTarget * .6 || offset < rightTarget * .6) {
      showConfrim();
    } else {
      reset();
    }
  }

  void onPanStart(e) {
    Utils.setTimeout(Duration(milliseconds: 0), () {
      // print(e);
      setState(() {
        target = e.localPosition.dx;
      });
    });
  }

  void onPanDown(e) {
    if (offset != 0) {
      reset();
    }
  }

  void onPanUpdate(e) {
    print("update:${e.localPosition.dx.toString()} target: $target ");
    var move = (e.localPosition.dx - target) * 0.75;
    setState(() {
      print(offset);
      // 正数向左滑动，负数向有滑动
      if (move > 0) {
        if (leftTarget == 0) return;
        if (leftOffset + move >= leftTarget && offset != 0) {
          offset = leftTarget;
          return;
        }
      } else {
        if (rightTarget == 0) return;
        if (rightOffset - move >= rightTarget && offset != 0) {
          offset = -rightTarget;
          return;
        }
      }

      offset += move;
      target = e.localPosition.dx;
      print(offset);
    });
  }

  Widget background(SlidingBackground child) {
    return GestureDetector(
      onPanDown: (e) {
        eventBus.fire(new SlidingEventsBus("inOperation"));
      },
      child: Container(
        height: widget.height,
        width: child.width,
        child: child.child,
      ),
    );
  }
}
