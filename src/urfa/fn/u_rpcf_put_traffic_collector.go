package fn

// rpcf_put_traffic_collector
func x5065(c conn, p Dict) Dict {
	c.Call(0x5065)
	putI(c, p, "id")
	putS(c, p, "name")
	putS(c, p, "comments")
	c.Send()

	return Dict{
		"error_message": c.GetS(),
	}
}
