package controllers

type Controller interface {
	Fetch(c interface{})
	One(c interface{})
	New(c interface{})
	Dup(c interface{})
	Edit(c interface{})
	Rename(c interface{})
	Delete(c interface{})
}
