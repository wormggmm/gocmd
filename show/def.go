package show

type IDrawable interface {
	ID() int64
	Draw()
	ApplyCursorPos()
	SetManager(manager *Manager)
}
