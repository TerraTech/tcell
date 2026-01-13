package tcell

// SetForceACS enables/disables forced ACS line drawing on screens that support it.
// It unwraps baseScreen wrappers. Returns true if applied.
func SetForceACS(s Screen, enable bool) bool {
	// Fast path: direct support.
	type forceACSSetter interface{ SetForceACS(bool) }
	if x, ok := s.(forceACSSetter); ok {
		x.SetForceACS(enable)
		return true
	}

	return false
}

func (b *baseScreen) SetForceACS(enable bool) {
	// Optional capability: only some underlying implementations (like *tScreen)
	// will support this.
	type forceACSSetter interface{ SetForceACS(bool) }
	if x, ok := b.screenImpl.(forceACSSetter); ok {
		x.SetForceACS(enable)
	}
}
