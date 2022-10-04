package fn

// rpcf_delete_shaping
func x12011(c conn, p Dict) Dict {
	c.Call(0x12011)
	putI(c, p, "service_id")
	putI(c, p, "delete_turbo_mode_settings")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
