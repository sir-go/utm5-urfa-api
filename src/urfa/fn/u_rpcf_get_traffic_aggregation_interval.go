package fn

// rpcf_get_traffic_aggregation_interval
func x10203(c conn, p Dict) Dict {
	c.Call(0x10203)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"aggregation_interval": c.GetI(),
	}
}
