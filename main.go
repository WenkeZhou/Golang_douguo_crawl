package main

import (
	"./engine"
	"./parse"
	"./persist"
	"./scheduler"
)

func main() {
	//engine.Run(engine.Request{
	//	Url: "https://www.douguo.com/fenlei",
	//	//Url: "https://www.douguo.com/caipu/%E8%99%BE%E4%BB%81",
	//	//Url: "https://www.douguo.com/cookbook/1588410.html",
	//	ParseFunc: parse.ParseContent,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSave(),
	}

	e.Run(engine.Request{
		Url:       "https://www.douguo.com/fenlei",
		ParseFunc: parse.ParseTag,
	})
}
