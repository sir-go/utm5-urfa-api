package fn

// rpcf_put_router
func x501d(c conn, p Dict) Dict {
	c.Call(0x501d)
	putI(c, p, "router_id", 0)
	putI(c, p, "router_type")
	putS(c, p, "router_ip")
	putS(c, p, "login")
	putS(c, p, "password")
	putS(c, p, "router_comments")
	putA(c, p, "router_bin_ip")
	c.Send()

	return nil
}
