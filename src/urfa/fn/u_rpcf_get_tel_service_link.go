package fn

// rpcf_get_tel_service_link
func x1033c(c conn, p Dict) Dict {
	c.Call(0x1033c)
	putI(c, p, "slinkId")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1033c" name="rpcf_get_tel_service_link">
    <input>
      <integer name="slinkId" />
    </input>
    <output>
      <integer name="accounting_period_id" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <integer name="policy_id" />
      <double name="cost_coef" />
      <integer name="house_id" />
      <string name="comment" />
      <integer name="numSize" />
      <for count="numSize" from="0" name="i">
	<string array_index="i" name="login" />
	<string array_index="i" name="number" />
	<string array_index="i" name="incoming_trunk" />
	<string array_index="i" name="outgoing_trunk" />
	<string array_index="i" name="pbx_id" />
	<string array_index="i" name="password" />
	<string array_index="i" name="allowed_cid" />
      </for>
      <integer name="slinkId" />
    </output>
  </function>


*/
