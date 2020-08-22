package main

import "strconv"

type ChaseLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *ChaseLoader) loadMasterPlaylist() ([]byte, error) {
	return []byte("ChaseLoader: loadMasterPlaylist"), nil
}

func (v *ChaseLoader) loadMediaPlaylist(id int) ([]byte, error) {
	return []byte("ChaseLoader: loadMediaPlaylist " + strconv.Itoa(id)), nil
}

func (v *ChaseLoader) loadSegment(id int) ([]byte, error) {
	return []byte("ChaseLoader: loadSegment " + strconv.Itoa(id)), nil
}
