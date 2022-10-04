package fn

// rpcf_get_sysgroups_list
func x2403(c conn, _ Dict) Dict {
	c.Call(0x2403)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2403" name="rpcf_get_sysgroups_list">
    <input />
    <output>
      <integer name="groups_size" />
      <for count="groups_size" from="0" name="i">
        <integer array_index="i" name="group_id" />
        <string array_index="i" name="group_name" />
        <string array_index="i" name="group_info" />
      </for>
    </output>
  </function>


*/
