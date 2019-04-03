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
			if v.Start.Nanos < sortedEvents[i+1].Start.Nanos {
				t.Errorf("Invalid sort expected %d to be less than %d", v.Start.Nanos, sortedEvents[i].Start.Nanos)
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
