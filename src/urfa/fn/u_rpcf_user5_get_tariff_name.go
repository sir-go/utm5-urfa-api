package fn

// rpcf_user5_get_tariff_name
func xU4039(c conn, p Dict) Dict {
	c.Call(-0x4039)
	putI(c, p, "tariff_id")
	c.Send()

	return Dict{
		"tariff_name": c.GetS(),
	}
}
