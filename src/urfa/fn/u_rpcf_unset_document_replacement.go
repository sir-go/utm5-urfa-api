package fn

// rpcf_unset_document_replacement
func x460c(c conn, p Dict) Dict {
	c.Call(0x460c)
	putS(c, p, "key")
	c.Send()

	return nil
}
