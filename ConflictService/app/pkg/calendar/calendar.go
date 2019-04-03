package calendar

import (
	"context"
	protobufs "scheduling"
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
	return &protobufs.ConflictList{}, nil
}
