package fn

// rpcf_get_tel_number_list
func x10319(c conn, _ Dict) Dict {
	c.Call(0x10319)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10319" name="rpcf_get_tel_number_list">
    <input>
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="num_id" />
        <integer array_index="i" name="slink_id" />
        <string array_index="i" name="login" />
        <string array_index="i" name="number" />
        <string array_index="i" name="incoming_trunk" />
        <string array_index="i" name="outgoing_trunk" />
        <string array_index="i" name="pbx_id" />
        <string array_index="i" name="password" />
        <string array_index="i" name="allowed_cid" />
      </for>
    </output>
  </function>



*/
