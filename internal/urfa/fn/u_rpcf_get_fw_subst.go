package fn

// rpcf_get_fw_subst
func x0049(c conn, _ Dict) Dict {
	c.Call(0x0049)

	return Dict{
		"items": getMapOf(c, func() Dict {
			return Dict{
				"subst":  c.GetS(),
				"events": c.GetL(),
			}
		}),
	}
}
