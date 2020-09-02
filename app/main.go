package main

import "database/sql"

func ¥¥

type sample2 struct {
	sampleInterface sampleInterface
}

func (s *sample2) A() {
	s.
}

func NewSampleStruct(db *sql.DB) sampleInterface {
	return &sampleStruct{db: db}
}

func NewMockSampleStruct() sampleInterface {
	return &mockSampleStruct{}
}

type sampleStruct struct {
	db *sql.DB
}

type mockSampleStruct struct{}

func (s *sampleStruct) sampleMethod() {
	s.db.Ping()
}

func (*mockSampleStruct) sampleMethod() {
}

type sampleInterface interface {
	sampleMethod()
}
