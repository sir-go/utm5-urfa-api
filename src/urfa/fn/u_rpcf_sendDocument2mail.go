package fn

// rpcf_sendDocument2mail
func x8005(c conn, p Dict) Dict {
	c.Call(0x8005)
	putI(c, p, "aid")
	putI(c, p, "invc_id")
	putI(c, p, "doc_type")
	c.Send()

	return nil
}
