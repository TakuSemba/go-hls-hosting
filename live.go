package main

import "strconv"

type LiveLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *LiveLoader) loadMasterPlaylist() ([]byte, error) {
	return []byte("LiveLoader: loadMasterPlaylist"), nil
}

func (v *LiveLoader) loadMediaPlaylist(index int) ([]byte, error) {
	return []byte("LiveLoader: loadMediaPlaylist " + strconv.Itoa(index)), nil
}

func (v *LiveLoader) loadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	return []byte("LiveLoader: loadSegment " + strconv.Itoa(segmentIndex)), nil
}
