package HqSMP

import "testing"

func TestMusicManager(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed , not empty")
	}
	mu0 := &MusicEntry{
		Id:"1",
		Name:"My Heart Will Go On",
		Artist:"Celion Dion",
		Source:"http://qbox.me/24501234",
		Type:"MP3",
	}

	mm.Add(mu0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed")
	}

	fm := mm.Find(mu0.Name)
	if fm == nil {
		t.Error("MusicManager.Find() failed")
	}
	if fm.Id != mu0.Id ||
		fm.Artist != mu0.Artist ||
		fm.Source != mu0.Source ||
		fm.Type != mu0.Type {
		t.Error("MusicManager.Find() failed.Fund item mismatch")
	}

	firstmu,err := mm.Get(0)
	if firstmu == nil{
		t.Error("MusicManager.Get() Failed",err)
	}
	rmmu := mm.Remove(0)
	if rmmu == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed",err)
	}
}