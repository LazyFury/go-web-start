import UIKit
import Flutter

@UIApplicationMain
@objc class AppDelegate: FlutterAppDelegate {
  override func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?
  ) -> Bool {
    

    
    GeneratedPluginRegistrant.register(with: self)
    
    let controller : FlutterViewController = window?.rootViewController as! FlutterViewController;
    let testFunc = FlutterMethodChannel.init(name: "samples.flutter.io/testFunc",
                                                   binaryMessenger: controller.binaryMessenger);
    testFunc.setMethodCallHandler({
      (call: FlutterMethodCall, result: FlutterResult) -> Void in
      // Handle battery messages.
    });
    
    
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}


private func testFunc (){
    print("ios testFunc");
}
