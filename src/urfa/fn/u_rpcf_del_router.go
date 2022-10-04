package fn

// rpcf_del_router
func x5012(c conn, p Dict) Dict {
	c.Call(0x5012)
	putI(c, p, "router_id")
	c.Send()

	return Dict{
		"router_code": c.GetI(),
	}
}
