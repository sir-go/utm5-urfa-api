package fn

// rpcf_edit_voluntary_suspension_settings
func x15012(c conn, p Dict) Dict {
	c.Call(0x15012)
	putI(c, p, "id")
	putI(c, p, "group_id")
	putI(c, p, "priopity")
	putI(c, p, "is_enabled")
	putI(c, p, "min_duration")
	putI(c, p, "max_duration")
	putI(c, p, "interval_duration")
	putI(c, p, "block_type")
	putI(c, p, "self_unlock")
	putD(c, p, "min_balance")
	putI(c, p, "use_min_balance")
	putD(c, p, "free_balance")
	putI(c, p, "use_free_balance")
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
