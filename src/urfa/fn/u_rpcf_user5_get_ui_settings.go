package fn

// rpcf_user5_get_ui_settings
func xU403b(c conn, _ Dict) Dict {
	c.Call(-0x403b)

	return Dict{
		"permissions": c.GetI(),
	}
}
