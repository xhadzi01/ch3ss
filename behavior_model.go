package main

type BehaviorModel struct {
}

func NewBehaviorModel() IBehaviorModel {
	return &BehaviorModel{}
}
