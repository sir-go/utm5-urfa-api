package fn

// rpcf_user5_get_tariff_id_by_name
func xU401a(c conn, p Dict) Dict {
	c.Call(-0x401a)
	putS(c, p, "name")
	c.Send()

	return Dict{
		"tid": c.GetI(),
	}
}
