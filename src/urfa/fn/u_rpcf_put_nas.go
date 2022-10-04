package fn

// rpcf_put_nas
func x5103(c conn, p Dict) Dict {
	c.Call(0x5103)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x5103" name="rpcf_put_nas">
    <input>

      <integer name="mode" />
      <integer name="id" />
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
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
