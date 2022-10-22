package fn

// rpcf_cportal_add_user
func x8102(c conn, p Dict) Dict {
	c.Call(0x8102)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"user_id":    c.GetI(),
		"account_id": c.GetI(),
	}
}

/* xml:
<function id="0x8102" name="rpcf_cportal_add_user">
    <input>
      <string name="login" />
      <string name="password" />
      <integer name="tariff_id" />
      <integer name="accounting_period_id" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <string array_index="i" name="parameter_name" />
        <string array_index="i" name="parameter_value" />
      </for>
    </input>
    <output>
      <integer name="user_id" />
      <integer name="account_id" />
    </output>
  </function>


*/
