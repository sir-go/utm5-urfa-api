package fn

// rpcf_remove_document_template
func x4604(c conn, p Dict) Dict {
	c.Call(0x4604)
	putI(c, p, "template_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
