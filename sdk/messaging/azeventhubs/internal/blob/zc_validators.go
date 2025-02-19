//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package blob

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Raw converts PageRange into primitive start, end integers of type int64
func (pr *PageRange) Raw() (start, end int64) {
	if pr.Start != nil {
		start = *pr.Start
	}
	if pr.End != nil {
		end = *pr.End
	}

	return
}

// HttpRange defines a range of bytes within an HTTP resource, starting at offset and
// ending at offset+count. A zero-value HttpRange indicates the entire resource. An HttpRange
// which has an offset but na zero value count indicates from the offset to the resource's end.
type HttpRange struct {
	Offset int64
	Count  int64
}

func NewHttpRange(offset, count int64) *HttpRange {
	return &HttpRange{Offset: offset, Count: count}
}

func (r *HttpRange) format() *string {
	if r == nil || (r.Offset == 0 && r.Count == 0) { // Do common case first for performance
		return nil // No specified range
	}
	endOffset := "" // if count == CountToEnd (0)
	if r.Count > 0 {
		endOffset = strconv.FormatInt((r.Offset+r.Count)-1, 10)
	}
	dataRange := fmt.Sprintf("bytes=%v-%s", r.Offset, endOffset)
	return &dataRange
}

func getSourceRange(offset, count *int64) *string {
	if offset == nil && count == nil {
		return nil
	}
	newOffset := int64(0)
	newCount := int64(CountToEnd)

	if offset != nil {
		newOffset = *offset
	}

	if count != nil {
		newCount = *count
	}

	return (&HttpRange{Offset: newOffset, Count: newCount}).format()
}

func validateSeekableStreamAt0AndGetCount(body io.ReadSeeker) (int64, error) {
	if body == nil { // nil body's are "logically" seekable to 0 and are 0 bytes long
		return 0, nil
	}

	err := validateSeekableStreamAt0(body)
	if err != nil {
		return 0, err
	}

	count, err := body.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, errors.New("body stream must be seekable")
	}

	_, err = body.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// return an error if body is not a valid seekable stream at 0
func validateSeekableStreamAt0(body io.ReadSeeker) error {
	if body == nil { // nil body's are "logically" seekable to 0
		return nil
	}
	if pos, err := body.Seek(0, io.SeekCurrent); pos != 0 || err != nil {
		// Help detect programmer error
		if err != nil {
			return errors.New("body stream must be seekable")
		}
		return errors.New("body stream must be set to position 0")
	}
	return nil
}
