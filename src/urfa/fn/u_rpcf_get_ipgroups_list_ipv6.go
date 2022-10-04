package fn

// rpcf_get_ipgroups_list_ipv6
func x292e(c conn, _ Dict) Dict {
	c.Call(0x292e)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x292e" name="rpcf_get_ipgroups_list_ipv6">
    <input />
    <output>
      <integer name="groups_size" />
      <for count="groups_size" from="0" name="i">
        <integer name="count" />
        <set dst="count_array" dst_index="i" src="count" />
        <for count="count" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <ip_address array_index="i,j" name="ip" />
          <integer array_index="i,j" name="mask" />
          <string array_index="i,j" name="mac" />
          <string array_index="i,j" name="login" />
          <string array_index="i,j" name="allowed_cid" />
        </for>
      </for>
    </output>
  </function>


*/
