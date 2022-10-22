package fn

// rpcf_get_tariff_id_by_name
func x301d(c conn, p Dict) Dict {
	c.Call(0x301d)
	putS(c, p, "name")
	c.Send()

	return Dict{
		"tid": c.GetI(),
	}
}
