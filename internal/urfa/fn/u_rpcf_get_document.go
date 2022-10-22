package fn

// rpcf_get_document
func x4620(c conn, p Dict) Dict {
	c.Call(0x4620)
	putI(c, p, "user_id")
	putI(c, p, "doc_type")
	putI(c, p, "base_id")
	putI(c, p, "file_type")
	c.Send()

	return Dict{
		"result":   c.GetI(),
		"doc_data": c.GetBin(),
	}
}
