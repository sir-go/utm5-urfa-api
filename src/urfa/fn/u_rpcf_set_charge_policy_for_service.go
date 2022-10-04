package fn

// rpcf_set_charge_policy_for_service
func x15106(c conn, p Dict) Dict {
	c.Call(0x15106)
	putI(c, p, "service_id")
	putI(c, p, "policy_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
