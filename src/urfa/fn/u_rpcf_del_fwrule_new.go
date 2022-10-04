package fn

// rpcf_del_fwrule_new
func x5023(c conn, p Dict) Dict {
	c.Call(0x5023)
	putI(c, p, "rule_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
