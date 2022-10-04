package fn

// enable_turbo_mode
func x293d(c conn, p Dict) Dict {
	c.Call(0x293d)
	putI(c, p, "slink_id")
	putI(c, p, "turbomode_settings_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
