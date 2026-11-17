package main

import (
	"fmt"
	// The import path is: module name (from go.mod) + folder.
	// Only names that start with a CAPITAL letter are visible from here.
	//"example.com/zinetic/week1-lessons/lesson0"
)

func main() {
	events := []lesson0.Event{
		{ID: "a", Result: "success"},
		{ID: "b", Result: "failure"},
		{ID: "c", Result: "success"},
	}

	fmt.Println("--- slice")
	ids := lesson0.SuccessfulIDs(events)
	fmt.Println("successful IDs:", ids, " count:", len(ids))

	fmt.Println("--- map")
	counts := lesson0.CountByResult(events)
	fmt.Println("counts:", counts)
	fmt.Println(`counts["timeout"]:`, counts["timeout"], " really present?", lesson0.HasResult(counts, "timeout"))

	fmt.Println("--- method")
	var c lesson0.Counter
	c.Add(2)
	c.Add(3)
	fmt.Println("total:", c.Total())

	fmt.Println("--- interface")
	fake := &lesson0.FakeSource{Pages: [][]lesson0.Event{events[:2], events[2:]}}
	fmt.Println("collected from 2 pages:", len(lesson0.CollectAll(fake)), "events")

	fmt.Println("--- errors")
	good := lesson0.ValidatePage(events, 10)
	fmt.Println("valid page    ->", good, "->", lesson0.Describe(good))
	missing := lesson0.ValidatePage([]lesson0.Event{{ID: "a"}, {ID: ""}}, 10)
	fmt.Println("empty ID      ->", missing, "->", lesson0.Describe(missing))
	tooMany := lesson0.ValidatePage(events, 2)
	fmt.Println("over the limit->", tooMany, "->", lesson0.Describe(tooMany))

	fmt.Println("--- defer")
	var ok, failed lesson0.Recorder
	_ = lesson0.DoWork(&ok, false)
	err := lesson0.DoWork(&failed, true)
	fmt.Println("normal path:", ok.Steps)
	fmt.Println("error path: ", failed.Steps, " error:", err)
}
