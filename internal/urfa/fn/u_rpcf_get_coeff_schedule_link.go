package fn

// rpcf_get_coeff_schedule_link
func x4710(c conn, p Dict) Dict {
	c.Call(0x4710)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":            c.GetI(),
		"scheme_id":     c.GetI(),
		"slink_id":      c.GetI(),
		"schedule_id":   c.GetI(),
		"change_policy": c.GetI(),
		"create_date":   c.GetI(),
		"change_date":   c.GetI(),
	}
}
