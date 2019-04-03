package calendar

import (
	protoTime "github.com/golang/protobuf/ptypes/timestamp"
	protobufs "scheduling"
	"testing"
	"time"
)

func TestSortEventList(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts now and ends in 10 minutes, append it
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	// Create an event that starts a 5 hours ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5 * time.Hour)),
		End:   TimestampProto(time.Now())})

	// Create an event that starts a 8 hours ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8 * time.Hour)),
		End:   TimestampProto(time.Now())})

	// Create an event that starts a 3 hours ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-3 * time.Hour)),
		End:   TimestampProto(time.Now())})

	// Create an event that starts a 2 hours ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-2 * time.Hour)),
		End:   TimestampProto(time.Now())})

	// Create an event that starts a 1 hour ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-1 * time.Hour)),
		End:   TimestampProto(time.Now())})

	// Sort them
	sortedEvents := sortEventList(testEvents)

	// This should be made better by actually comparing a full list
	for i, v := range sortedEvents {
		if i != len(sortedEvents)-1 {
			if v.Start.Seconds > sortedEvents[i+1].Start.Seconds {
				t.Errorf("Invalid sort expected %d to be less than %d", v.Start.Seconds, sortedEvents[i].Start.Seconds)
			}
		}
	}
}

// Helper function for the unit test
func TimestampProto(t time.Time) *protoTime.Timestamp {
	ts := &protoTime.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}

	return ts
}

func TestGetConflictGroupsNoConflicts(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts now and ends in 10 minutes, append it
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	eventList := &protobufs.EventList{Events: testEvents}

	conflictList := getConflictGroups(eventList)

	got := len(conflictList.Conflicts)

	if got > 0 {
		t.Errorf("Invalid conflict detection expected none, got %d", got)
	}
}

func TestGetConflictGroupsSimpleConflictAtEnd(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts now and ends in 10 minutes, append it
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	// Create an event that starts a year ago and ends 2 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8758 * time.Hour))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	eventList := &protobufs.EventList{Events: testEvents}

	conflictList := getConflictGroups(eventList)

	got := len(conflictList.Conflicts)
	got1 := len(conflictList.Conflicts[0].ConflictGroup)

	if got != 1 || got1 != 2 {
		t.Errorf("Invalid conflict detection expected 1 group wtih 2 in it, got %d groups and %d in group 1", got, got1)
	}
}

//Make sure we handle cases where 3 calendar events conflict
func TestGetConflictGroupsMultipleConflicts(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts now and ends in 10 minutes, append it
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	// Create an event that starts a year ago and ends 2 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8758 * time.Hour))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	eventList := &protobufs.EventList{Events: testEvents}

	conflictList := getConflictGroups(eventList)

	got := len(conflictList.Conflicts)
	got1 := len(conflictList.Conflicts[0].ConflictGroup)

	if got1 != 3 || got != 1 {
		t.Errorf("Invalid conflict detection expected 3 in 1 group, got %d in %d groups", got1, got)
	}
}

//Make sure we handle cases where we have multiple instances of 3+ calendar events conflicting
func TestGetConflictGroupsMultipleConflictsMultipleGroups(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts now and ends in 10 minutes, append it
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	// Create an event that starts a year ago and ends 2 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8758 * time.Hour))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	// Create an event that starts a year ago and ends 1 hour later
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	// Create some random middle events that all conflict with eachother but nothing else
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5758 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5759 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5759 * time.Hour))})

	eventList := &protobufs.EventList{Events: testEvents}

	conflictList := getConflictGroups(eventList)

	got := len(conflictList.Conflicts)
	got1 := len(conflictList.Conflicts[0].ConflictGroup)
	got2 := len(conflictList.Conflicts[1].ConflictGroup)

	if got != 2 || got1 != 3 || got2 != 3 {
		t.Errorf("Invalid conflict detection expected 3 in 2 groups, got %d groups and %d in group 1 and %d in group 2", got, got1, got2)
	}
}

// Create a really long conflicting event that will conflict with everything
func TestGetConflictGroupsEventConflictsEverything(t *testing.T) {
	// Create an event list
	var testEvents []*protobufs.Event

	// Create an event that starts a year ago and ends now
	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now())})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8758 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-8760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-8759 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5758 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5759 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now().Add(-5760 * time.Hour)),
		End:   TimestampProto(time.Now().Add(-5759 * time.Hour))})

	testEvents = append(testEvents, &protobufs.Event{
		Start: TimestampProto(time.Now()),
		End:   TimestampProto(time.Now().Add(10 * time.Minute))})

	eventList := &protobufs.EventList{Events: testEvents}

	conflictList := getConflictGroups(eventList)

	got := len(conflictList.Conflicts)
	got1 := len(conflictList.Conflicts[0].ConflictGroup)

	if got != 1 || got1 != 7 {
		t.Errorf("Invalid conflict detection expected 7 in 1 group, got %d groups and %d in group 1", got, got1)
	}
}
