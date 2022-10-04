package fn

// rpcf_user5_activate_megogo_tariff
func xU15041(c conn, p Dict) Dict {
	c.Call(-0x15041)
	putI(c, p, "tariff_plan_id")
	c.Send()

	return Dict{
		"status": c.GetI(),
	}
}
