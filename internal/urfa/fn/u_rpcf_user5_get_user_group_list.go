package fn

// rpcf_user5_get_user_group_list
func xU401c(c conn, _ Dict) Dict {
	c.Call(-0x401c)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x401c" name="rpcf_user5_get_user_group_list">
    <input />
    <output>
      <integer name="groups_count" />
      <for count="groups_count" from="0" name="i">
        <integer array_index="i" name="id" />
      </for>
    </output>
  </function>


*/
