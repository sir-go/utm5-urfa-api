package fn

// rpcf_del_service_from_tariff
func x3015(c conn, p Dict) Dict {
	c.Call(0x3015)
	putI(c, p, "tp_id")
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
