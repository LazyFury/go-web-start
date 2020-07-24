package chat

// Group Group
type Group map[string]*Gamer

// Group
func (g Group) hasKey(id string) (hasKey bool) {
	_, hasKey = g[id]
	return
}

func (g Group) remove(id string) (user *Gamer, ok bool) {
	user, ok = g[id]
	delete(g, id)
	return user, ok
}
