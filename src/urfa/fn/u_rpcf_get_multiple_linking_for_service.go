package fn

// rpcf_get_multiple_linking_for_service
func x1510f(c conn, p Dict) Dict {
	c.Call(0x1510f)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"multiple_linking": c.GetI(),
	}
}
