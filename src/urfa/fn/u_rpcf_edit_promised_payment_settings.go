package fn

// rpcf_edit_promised_payment_settings
func x15036(c conn, p Dict) Dict {
	c.Call(0x15036)
	putI(c, p, "id")
	putI(c, p, "group_id")
	putI(c, p, "priopity")
	putI(c, p, "is_enabled")
	putD(c, p, "max_value")
	putI(c, p, "max_duration")
	putI(c, p, "interval_duration")
	putD(c, p, "min_balance")
	putI(c, p, "use_min_balance")
	putD(c, p, "free_balance")
	putI(c, p, "use_free_balance")
	putI(c, p, "service_id")
	putI(c, p, "flags")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
