package main

import "fmt"

type Boss struct{}

func (b *Boss) AssignWork() {
	fmt.Println("Boss assigned work")
}

type Manager struct{}

func (m *Manager) PreparePowerPoint() {
	fmt.Println("PowerPoint prepared")
}

type BossManager struct {
	Boss
	Manager
}

type BossManagerUsingPointers struct {
	*Boss
	*Manager
}

type PromotionMaterial interface {
	AssignWork()
	PreparePowerPoint()
}

func promote(pm PromotionMaterial) {
	fmt.Println("Promoted a person with promise.")
}

func main() {
	bm := BossManager{}
	bm.PreparePowerPoint()
	bm.AssignWork()
	promote(&bm)

	bm2 := BossManagerUsingPointers{}
	promote(bm2)

}
