package calendar

import (
	"context"
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

	conflictGroups := getConflictGroups(eventList)

	return conflictGroups, nil
}

func getConflictGroups(eventList *protobufs.EventList) *protobufs.ConflictList {
	// First step, sort them in nlogn time by start time
	events := sortEventList(eventList.Events)
	numEvents := len(events)
	conflictList := &protobufs.ConflictList{}

	if numEvents > 0 {
		// Prime everything because we are starting to look at the second item in the scheduling list
		endTime := events[0].End.Nanos
		tmpConflictGroup := &protobufs.ConflictGroup{ConflictGroup: make([]*protobufs.Event, 0)}
		tmpConflictGroup.ConflictGroup = append(tmpConflictGroup.ConflictGroup, events[0])

		for i := 1; i < numEvents; i++ {
			// Is the start time of the current event, after the end time of the previous longest ending event
			if events[i].Start.Nanos >= endTime {
				// If it is, then we do not have a conflict
				// Therefore we can add the previous conflict group to the conflict list, if it has any scheduling conflicts in it
				if len(tmpConflictGroup.ConflictGroup) > 1 {
					// Also worth noting, you need at least two events to have a conflict
					conflictList.Conflicts = append(conflictList.Conflicts, tmpConflictGroup)
				}

				// And we can flush the current buffer of conflicting events
				tmpConflictGroup = &protobufs.ConflictGroup{ConflictGroup: make([]*protobufs.Event, 0)}
			}

			// Make sure the end time is always the latest end time, this will handle cases where event B ends before event A but event A still conflicts with event C
			endTime = Max(endTime, events[i].End.Nanos)
			tmpConflictGroup.ConflictGroup = append(tmpConflictGroup.ConflictGroup, events[i])
		}

		// Handle the case where the very last event in the loop conflicts
		if len(tmpConflictGroup.ConflictGroup) > 1 {
			conflictList.Conflicts = append(conflictList.Conflicts, tmpConflictGroup)
		}
	}

	return &protobufs.ConflictList{}
}

func Max(a, b int32) int32 {
	if a < b {
		return b
	}

	return a
}

func sortEventList(events []*protobufs.Event) []*protobufs.Event {
	// Go's built in sort, sorts in nLogn time..
	sort.Slice(events, func(a, b int) bool {
		return events[a].Start.Nanos > events[b].Start.Nanos
	})

	return events
}
