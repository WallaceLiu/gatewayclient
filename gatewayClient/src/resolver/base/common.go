package base

import "strings"

//
func FindCollectionSegmentFirstById(id string, segment []Segment) *Segment {
	for _, d := range segment {

		if strings.ToLower(id) == strings.ToLower(d.Id) {
			return &d
		} else {
			if d.Segment != nil && len(d.Segment) > 0 {
				return findSegmentCollectionFirstById(id, d)
			}
		}
	}

	return nil
}

//
func findSegmentCollectionFirstById(id string, segment Segment) *Segment {
	for _, d := range segment.Segment {

		if strings.ToLower(id) == strings.ToLower(d.Id) {
			return &d
		} else {
			if d.Segment != nil && len(d.Segment) > 0 {
				return findSegmentCollectionFirstById(id, d)
			}
		}
	}

	return nil
}
