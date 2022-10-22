package fn

// rpcf_cportal_get_daily_accouting_period
func x8100(c conn, p Dict) Dict {
	c.Call(0x8100)
	putI(c, p, "periodic_type")
	putI(c, p, "custom_duration")
	c.Send()

	return Dict{
		"period_id": c.GetI(),
	}
}
