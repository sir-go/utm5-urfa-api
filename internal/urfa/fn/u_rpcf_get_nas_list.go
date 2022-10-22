package fn

// rpcf_get_nas_list
func x5040(c conn, _ Dict) Dict {
	c.Call(0x5040)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5040" name="rpcf_get_nas_list">
    <input />
    <output>
      <integer name="nas_size" />
      <for count="nas_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="nas_id" />
        <string array_index="i" name="auth_secret" />
        <string array_index="i" name="acct_secret" />
        <string array_index="i" name="dac_secret" />
        <integer array_index="i" name="das_port" />
        <integer array_index="i" name="flags" />
      </for>
    </output>
  </function>


*/
