package fn

// rpcf_user5_get_service_id_by_name
func xU401e(c conn, p Dict) Dict {
	c.Call(-0x401e)
	putS(c, p, "name")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
