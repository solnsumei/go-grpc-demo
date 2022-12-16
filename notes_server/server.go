package main

import (
	"context"
	"log"

	"github.com/solnsumei/go-grpc-demo/notes"
)

type NotesServer struct {
	notes.UnimplementedNotesServer
}

func (s *NotesServer) Save(ctx context.Context, n *notes.Note) (*notes.NoteSaveReply, error) {
	log.Printf("Received a note to save: %v\n", n.Title)
	err := notes.SaveToDisk(n, "testdata")

	if err != nil {
		return &notes.NoteSaveReply{Saved: false}, err
	}

	return &notes.NoteSaveReply{Saved: true}, nil
}

func (s *NotesServer) Load(ctx context.Context, search *notes.NoteSearch) (*notes.Note, error) {
	log.Printf("Received a note to load: %v\n", search.Keyword)
	note, err := notes.LoadFromDisk(search.Keyword, "testdata")
	if err != nil {
		return &notes.Note{}, err
	}

	return note, nil
}
