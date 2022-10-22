package fn

// rpcf_add_fwrule_new
func x5021(c conn, p Dict) Dict {
	c.Call(0x5021)
	putI(c, p, "flags")
	putL(c, p, "events")
	putI(c, p, "router_id")
	putI(c, p, "tariff_id")
	putI(c, p, "group_id")
	putI(c, p, "user_id")
	putS(c, p, "rule")
	putS(c, p, "comment")
	c.Send()

	return Dict{
		"rule_id": c.GetI(),
	}
}
