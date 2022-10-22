package fn

// rpcf_get_groups_for_user
func x2550(c conn, p Dict) Dict {
	c.Call(0x2550)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2550" name="rpcf_get_groups_for_user">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="groups_size" />
      <for count="groups_size" from="0" name="i">
        <integer array_index="i" name="group_id" />
        <string array_index="i" name="group_name" />
      </for>
    </output>
  </function>


*/
