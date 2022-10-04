package fn

// rpcf_get_version_info
func x5013(c conn, _ Dict) Dict {
	c.Call(0x5013)

	return Dict{
		"name":       c.GetS(),
		"country":    c.GetS(),
		"region":     c.GetS(),
		"city":       c.GetS(),
		"address":    c.GetS(),
		"email":      c.GetS(),
		"tel":        c.GetS(),
		"web":        c.GetS(),
		"valid_till": c.GetS(),
	}
}
