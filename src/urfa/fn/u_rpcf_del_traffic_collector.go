package fn

// rpcf_del_traffic_collector
func x5066(c conn, p Dict) Dict {
	c.Call(0x5066)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"error_message": c.GetS(),
	}
}
