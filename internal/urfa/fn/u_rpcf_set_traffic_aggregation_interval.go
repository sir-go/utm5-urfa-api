package fn

// rpcf_set_traffic_aggregation_interval
func x10204(c conn, p Dict) Dict {
	c.Call(0x10204)
	putI(c, p, "service_id")
	putI(c, p, "aggregation_interval")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
