package main

import (
	"golvgl/lvgl"
	"log"
	"time"
)

func main() {
	log.Println("Hello Go,LittlevGL")
	lb := lvgl.Label(lvgl.ScrAct(), nil)
	lb.SetText("hello world,go test 222")
	lb.Align(nil, lvgl.LV_ALIGN_CENTER, 0, 0)
	log.Println("Hello Go Over 222")
	for true {
		lvgl.TickInc(5)
		lvgl.TaskHandler()
		time.Sleep(5000)
	}
}
