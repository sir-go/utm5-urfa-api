package fn

// rpcf_user5_enable_turbo_mode
func xU1200d(c conn, p Dict) Dict {
	c.Call(-0x1200d)
	putI(c, p, "slink_id")
	putI(c, p, "turbomode_settings_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
