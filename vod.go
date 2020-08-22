package main

import "strconv"

type VodLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *VodLoader) loadMasterPlaylist() ([]byte, error) {
	return []byte("VodLoader: loadMasterPlaylist"), nil
}

func (v *VodLoader) loadMediaPlaylist(id int) ([]byte, error) {
	return []byte("VodLoader: loadMediaPlaylist " + strconv.Itoa(id)), nil
}

func (v *VodLoader) loadSegment(id int) ([]byte, error) {
	return []byte("VodLoader: loadSegment " + strconv.Itoa(id)), nil
}
