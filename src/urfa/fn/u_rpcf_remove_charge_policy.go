package fn

// rpcf_remove_charge_policy
func x15105(c conn, p Dict) Dict {
	c.Call(0x15105)
	putI(c, p, "policy_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
