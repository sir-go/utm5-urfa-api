package fn

// rpcf_edit_coeff_schedule_link
func x4713(c conn, p Dict) Dict {
	c.Call(0x4713)
	putI(c, p, "id")
	putI(c, p, "scheme_id")
	putI(c, p, "slink_id")
	putI(c, p, "schedule_id")
	putI(c, p, "change_policy")
	putI(c, p, "create_date")
	putI(c, p, "change_date")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
