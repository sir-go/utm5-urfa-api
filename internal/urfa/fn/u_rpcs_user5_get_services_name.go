package fn

// rpcs_user5_get_services_name
func xU402b(c conn, p Dict) Dict {
	c.Call(-0x402b)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"service_type":    c.GetI(),
		"service_id":      c.GetI(),
		"service_name":    c.GetS(),
		"service_comment": c.GetS(),
		"periodic_cost":   c.GetD(),
	}
}
