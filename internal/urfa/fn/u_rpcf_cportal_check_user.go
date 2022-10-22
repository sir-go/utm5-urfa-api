package fn

// rpcf_cportal_check_user
func x8101(c conn, p Dict) Dict {
	c.Call(0x8101)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"err_code":    c.GetI(),
		"err_message": c.GetS(),
	}
}

/* xml:
<function id="0x8101" name="rpcf_cportal_check_user">
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
      <integer name="err_code" />
      <string name="err_message" />
    </output>
  </function>


*/
