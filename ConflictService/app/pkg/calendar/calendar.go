package calendar

import (
	"context"
	"fmt"
	protobufs "scheduling"
	"sort"
)

// Server represents our implementation of the scheduling interface
type Server struct {
}

// NewServer -- (): Create a new server that implements the ConflictService rpc server interface
func NewServer() *Server {
	return &Server{}
}

// GetConflicts -- (ctx, conflictList) Returns a sequence of events that conflict with eachother
func (s *Server) GetConflicts(ctx context.Context, eventList *protobufs.EventList) (*protobufs.ConflictList, error) {

	conflictPairs := getConflictPairs(eventList)

	return conflictPairs, nil
}

func getConflictPairs(eventList *protobufs.EventList) *protobufs.ConflictList {
	//First step, sort them in nlogn time by start time
	_ = sortEventList(eventList)

	//Next step, just iterate down the events checking to see if the end time of n is at or after the start time of n + 1

	return &protobufs.ConflictList{}
}

func sortEventList(eventList *protobufs.EventList) []*protobufs.Event {
	events := eventList.Events

	// Go's built in sort, sort's in nLogn time..
	sort.Slice(events, func(a, b int) bool {
		return events[a].Start.Nanos < events[b].Start.Nanos
	})

	fmt.Println(events)

	return events
}
