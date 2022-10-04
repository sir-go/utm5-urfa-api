package fn

// rpcf_get_archive_list
func x00fd(c conn, _ Dict) Dict {
	c.Call(0x00fd)

	return Dict{
		"items": getMapOf(c, func() Dict {
			return Dict{
				"id":         c.GetI(),
				"begin_date": c.GetS(),
				"end_date":   c.GetS(),
				"status":     c.GetI(),
			}
		}),
	}
}
