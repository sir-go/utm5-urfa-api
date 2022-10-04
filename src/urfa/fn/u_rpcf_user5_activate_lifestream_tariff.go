package fn

// rpcf_user5_activate_lifestream_tariff
func xU15051(c conn, p Dict) Dict {
	c.Call(-0x15051)
	putI(c, p, "tariff_plan_id")
	c.Send()

	return Dict{
		"status": c.GetI(),
	}
}
