package main

import "strconv"

type LiveLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *LiveLoader) loadMasterPlaylist() ([]byte, error) {
	return []byte("LiveLoader: loadMasterPlaylist"), nil
}

func (v *LiveLoader) loadMediaPlaylist(id int) ([]byte, error) {
	return []byte("LiveLoader: loadMediaPlaylist " + strconv.Itoa(id)), nil
}

func (v *LiveLoader) loadSegment(id int) ([]byte, error) {
	return []byte("LiveLoader: loadSegment " + strconv.Itoa(id)), nil
}
