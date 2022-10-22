package fn

// rpcf_get_nas
func x5102(c conn, p Dict) Dict {
	c.Call(0x5102)
	putI(c, p, "id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5102" name="rpcf_get_nas">
    <input>
      <integer name="id" />
    </input>
    <output>
      <string name="nas_id" />
      <string name="auth_secret" />
      <string name="acct_secret" />
      <string name="dac_secret" />
      <integer name="das_port" />
      <integer name="flags" />
      <integer name="isg_profile" />
      <integer name="parameters_size" />
      <for count="parameters_size" from="0" name="i">
        <string array_index="i" name="variable" />
        <string array_index="i" name="value" />
      </for>
    </output>
  </function>


*/
