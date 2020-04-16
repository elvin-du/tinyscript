package util

import (
	"bufio"
	"container/list"
	"io"
)

type Stream struct {
	scanner    *bufio.Scanner
	queueCache *list.List
	endToken   string
	isEnd      bool
}

func NewStream(r io.Reader, et string) *Stream {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	return &Stream{scanner: s, queueCache: list.New(), endToken: et, isEnd: false}
}

func (s *Stream) Next() string {
	if s.queueCache.Len() != 0 {
		e := s.queueCache.Front()
		return s.queueCache.Remove(e).(string)
	}

	if s.scanner.Scan() {
		return s.scanner.Text()
	}

	s.isEnd = true

	return s.endToken
}

func (s *Stream) HasNext() bool {
	if s.queueCache.Len() != 0 {
		return true
	}

	if s.scanner.Scan() {
		s.queueCache.PushBack(s.scanner.Text())
		return true
	}

	if !s.isEnd {
		return true
	}

	return false
}

func (s *Stream) Peek() string {
	if s.queueCache.Len() != 0 {
		return s.queueCache.Front().Value.(string)
	}

	if s.scanner.Scan() {
		e := s.scanner.Text()
		s.queueCache.PushBack(e)
		return e
	}

	return s.endToken
}

func (s *Stream) PutBack(e string) {
	s.queueCache.PushFront(e)
}
