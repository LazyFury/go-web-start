//
//  CalendarManager.m
//  test
//
//  Created by suke on 2020/2/20.
//  Copyright © 2020 Facebook. All rights reserved.
//

#import "CalendarManager.h"
#import <React/RCTLog.h>

@implementation CalendarManager

// To export a module named CalendarManager
RCT_EXPORT_MODULE();

// This would name the module AwesomeCalendarManager instead
// RCT_EXPORT_MODULE(AwesomeCalendarManager);
RCT_EXPORT_METHOD(addEvent:(NSString *)name location:(NSString *)location)
{
  RCTLog(@"Pretending to create an event %@ at %@", name, location);
}

@end
