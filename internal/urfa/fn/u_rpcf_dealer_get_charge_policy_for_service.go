package fn

// rpcf_dealer_get_charge_policy_for_service
func x70000067(c conn, p Dict) Dict {
	c.Call(0x70000067)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"charge_policy": c.GetI(),
	}
}
