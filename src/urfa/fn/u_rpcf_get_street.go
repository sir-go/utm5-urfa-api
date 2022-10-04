package fn

// rpcf_get_street
func x2818(c conn, p Dict) Dict {
	c.Call(0x2818)
	putI(c, p, "street_id")
	c.Send()

	return Dict{
		"street_id": c.GetI(),
		"country":   c.GetS(),
		"region":    c.GetS(),
		"city":      c.GetS(),
		"street":    c.GetS(),
	}
}
