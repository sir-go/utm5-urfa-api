package fn

// rpcf_get_group_info
func x2409(c conn, p Dict) Dict {
	c.Call(0x2409)
	putI(c, p, "group_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2409" name="rpcf_get_group_info">
    <input>
      <integer name="group_id" />
    </input>
    <output>
      <string name="group_name" />
      <integer name="user_id_size" />
      <for count="user_id_size" from="0" name="i">
        <integer array_index="i" name="user_id" />
        <string array_index="i" name="login" />
      </for>
    </output>
  </function>


*/
