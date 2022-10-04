package fn

// rpcf_set_multiple_linking_for_service
func x15110(c conn, p Dict) Dict {
	c.Call(0x15110)
	putI(c, p, "service_id")
	putI(c, p, "multiple_linking")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
