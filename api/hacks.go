package api

func (f GetConnectionRow) Status() string {
	if f.RoleOut == "" {
		if f.RoleIn == "" {
			return "none"
		} else {
			return "request received"
		}
	} else {
		if f.RoleIn == "" {
			return "request sent"
		} else {
			return "connected"
		}
	}
}

func (f GetCurrentAndPotentialParentConnectionsRow) Status() string {
	if f.RoleOut == "" {
		if f.RoleIn == "" {
			return "none"
		} else {
			return "request received"
		}
	} else {
		if f.RoleIn == "" {
			return "request sent"
		} else {
			return "connected"
		}
	}
}
