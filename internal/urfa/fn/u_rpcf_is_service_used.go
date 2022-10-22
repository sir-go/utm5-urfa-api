package fn

// rpcf_is_service_used
func x10001(c conn, p Dict) Dict {
	c.Call(0x10001)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"links_count": c.GetI(),
	}
}
