package fn

// rpcf_user5_get_turbo_mode_settings
func xU1200b(c conn, p Dict) Dict {
	c.Call(-0x1200b)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"incoming_rate": c.GetI(),
		"outgoing_rate": c.GetI(),
		"duration":      c.GetI(),
		"cost":          c.GetD(),
	}
}
