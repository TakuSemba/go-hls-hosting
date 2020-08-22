package main

import "strconv"

type ChaseLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *ChaseLoader) loadMasterPlaylist() ([]byte, error) {
	return []byte("ChaseLoader: loadMasterPlaylist"), nil
}

func (v *ChaseLoader) loadMediaPlaylist(index int) ([]byte, error) {
	return []byte("ChaseLoader: loadMediaPlaylist " + strconv.Itoa(index)), nil
}

func (v *ChaseLoader) loadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	return []byte("ChaseLoader: loadSegment " + strconv.Itoa(segmentIndex)), nil
}
