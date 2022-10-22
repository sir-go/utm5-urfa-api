package fn

// rpcf_set_document_replacement
func x460b(c conn, p Dict) Dict {
	c.Call(0x460b)
	putS(c, p, "key")
	putS(c, p, "value")
	c.Send()

	return nil
}
