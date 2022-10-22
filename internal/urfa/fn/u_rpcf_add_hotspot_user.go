package fn

// rpcf_add_hotspot_user
func x4292(c conn, p Dict) Dict {
	c.Call(0x4292)
	putS(c, p, "tel_number", "8 800 000 00 00")
	putI(c, p, "card_pool_id", 0)
	putI(c, p, "tariff_id", 0)
	putA(c, p, "client_ip", "127.0.0.1")
	putI(c, p, "till", 0)
	putS(c, p, "comments")
	putS(c, p, "password")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4292" name="rpcf_add_hotspot_user">
    <input>
      <string default="8 800 000 00 00" name="tel_number" />
      <integer default="0" name="card_pool_id" />
      <integer default="0" name="tariff_id" />
      <ip_address default="127.0.0.1" name="client_ip" />
      <integer default="0" name="till" />
      <string default="" name="comments" />
      <string default="" name="password" />
    </input>
    <output>
      <integer name="user_id" />
        <if condition="ne" value="0" variable="user_id">
            <integer name="account_id" />
            <string name="login" />
            <string name="password" />
        </if>
        <if condition="eq" value="0" variable="user_id">
            <string name="error" />
        </if>
    </output>
  </function>








*/
