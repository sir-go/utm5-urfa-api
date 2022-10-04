package fn

// rpcf_edit_uaparam
func x440e(c conn, p Dict) Dict {
	c.Call(0x440e)
	putI(c, p, "id")
	putS(c, p, "name")
	putS(c, p, "display_name")
	putI(c, p, "visible")
	c.Send()

	return nil
}
