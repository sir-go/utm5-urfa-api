package fn

// rpcf_get_techparam_slink_by_uid
func x9003(c conn, p Dict) Dict {
	c.Call(0x9003)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9003" name="rpcf_get_techparam_slink_by_uid">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="slink_id" />
        <string array_index="i" name="service_name" />
      </for>
    </output>
  </function>


*/
