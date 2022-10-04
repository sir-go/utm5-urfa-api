package fn

// rpcf_put_tel_call
func x10310(c conn, p Dict) Dict {
	c.Call(0x10310)
	putS(c, p, "calling_number")
	putS(c, p, "called_number")
	putS(c, p, "incoming_trunk")
	putS(c, p, "outgoing_trunk")
	putS(c, p, "session_id")
	putS(c, p, "pbx_id")
	putS(c, p, "login")
	putS(c, p, "disconnect_cause")
	putI(c, p, "call_start_date", 0)
	putI(c, p, "call_duration", 0)
	c.Send()

	return nil
}
