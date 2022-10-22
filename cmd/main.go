/*
Copyright 2022 The Workpieces LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/syncthing/notify"
	"log"
)

func main() {
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening on events within current working directory.
	// Dispatch each create and remove events separately to c.
	// Y:\1006100210002\Info\...
	// 1. 测试 （Y:\1006100210002\Info\...）
	// 2. /Users/taoshumin_vendor/go/src/yho.io/notify/cmd/...
	if err := notify.Watch("Y:\\1006100210002\\Info\\...", c, notify.All); err != nil {
		//if err := notify.Watch("/Users/taoshumin_vendor/go/src/yho.io/notify/cmd/...", c, notify.All); err != nil {
		log.Fatal(err)
	}

	defer notify.Stop(c)

	for {
		select {
		case ei := <-c:
			log.Println("Got event:", ei)

		}
	}
}
