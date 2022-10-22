package fn

// rpcf_get_charge_policy_for_service
func x15107(c conn, p Dict) Dict {
	c.Call(0x15107)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"charge_policy": c.GetI(),
	}
}
