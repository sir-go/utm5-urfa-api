package fn

// rpcf_get_sys_settings
func x2035(c conn, _ Dict) Dict {
	c.Call(0x2035)

	return Dict{
		"block_recalc_abon":    c.GetI(),
		"block_recalc_prepaid": c.GetI(),
		"default_vat_rate":     c.GetD(),
	}
}
