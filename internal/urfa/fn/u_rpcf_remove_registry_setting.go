package fn

// rpcf_remove_registry_setting
func x5072(c conn, p Dict) Dict {
	c.Call(0x5072)
	putS(c, p, "name")
	putI(c, p, "object_id")
	c.Send()

	return nil
}
