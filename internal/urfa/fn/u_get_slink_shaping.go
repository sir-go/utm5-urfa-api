package fn

// get_slink_shaping
func x12009(c conn, p Dict) Dict {
	c.Call(0x12009)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"result":                    c.GetI(),
		"flags":                     c.GetI(),
		"turbomode_settings_id":     c.GetI(),
		"incoming_rate":             c.GetI(),
		"outgoing_rate":             c.GetI(),
		"turbo_mode_start":          c.GetI(),
		"turbo_mode_end":            c.GetI(),
		"incoming_consumption_left": c.GetL(),
		"outgoing_consumption_left": c.GetL(),
	}
}
