package fn

// rpcf_delete_switch_tariff_settings
func x15003(c conn, p Dict) Dict {
	c.Call(0x15003)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
