package fn

// rpcf_core_version_user
func xU0045(c conn, _ Dict) Dict {
	c.Call(-0x0045)

	return Dict{
		"version": c.GetS(),
	}
}
