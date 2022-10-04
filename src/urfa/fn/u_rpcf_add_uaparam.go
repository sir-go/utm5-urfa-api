package fn

// rpcf_add_uaparam
func x440c(c conn, p Dict) Dict {
	c.Call(0x440c)
	putI(c, p, "id")
	putS(c, p, "name")
	putS(c, p, "display_name")
	putI(c, p, "visible")
	c.Send()

	return nil
}
