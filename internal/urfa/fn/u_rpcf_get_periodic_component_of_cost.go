package fn

// rpcf_get_periodic_component_of_cost
func x10000(c conn, p Dict) Dict {
	c.Call(0x10000)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"cost": c.GetD(),
	}
}
