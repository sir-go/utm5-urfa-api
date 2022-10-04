package fn

// rpcf_add_discount_period
func x2603(c conn, p Dict) Dict {
	c.Call(0x2603)
	putI(c, p, "id")
	putI(c, p, "start")
	putI(c, p, "expire")
	putI(c, p, "periodic_type_t")
	putI(c, p, "cd")
	putI(c, p, "di")
	c.Send()

	return nil
}
