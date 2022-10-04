package fn

// rpcf_remove_document
func x4622(c conn, p Dict) Dict {
	c.Call(0x4622)
	putI(c, p, "user_id")
	putI(c, p, "doc_type")
	putI(c, p, "base_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
