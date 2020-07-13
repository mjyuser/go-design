package fly_weight

import "fmt"

// 享元模式 - 剥离公共元素
type TreeType struct {
	Name string
	Color string
	Texture string
}

func (t *TreeType) Draw (canvas Canvas, x, y int) {}

type Canvas struct {}

type Draw interface {
	Draw(canvas Canvas, x, y int)
}


var treeFactory = &TreeFactory{
	TreeTypes: make(map[string]*TreeType),
}

type TreeFactory struct {
	TreeTypes map[string]*TreeType
}

func (t *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := fmt.Sprintf("%s:%s:%s", name, color, texture)
	if tree, ok := t.TreeTypes[key]; ok {
		return tree
	}

	tree := &TreeType{
		Name:    name,
		Color:   color,
		Texture: texture,
	}

	t.TreeTypes[key] = tree
	return tree
}

type Tree struct {
	X,Y int
	tree *TreeType
}

func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		X:    x,
		Y:    y,
		tree: treeType,
	}
}

type Forest struct {
	trees []*Tree
}

func (f *Forest) PlantTree(x, y int, color, name, texture string) {
	treeType := treeFactory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw(canvas Canvas) {
	for _, tree := range f.trees {
		tree.tree.Draw(canvas, tree.X, tree.Y)
	}
}
