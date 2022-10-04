package fn

// rpcf_revert_document
func x4623(c conn, p Dict) Dict {
	c.Call(0x4623)
	putI(c, p, "user_id")
	putI(c, p, "doc_type")
	putI(c, p, "base_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
