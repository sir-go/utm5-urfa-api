package fn

// rpcf_get_sys_users_list
func x4413(c conn, _ Dict) Dict {
	c.Call(0x4413)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4413" name="rpcf_get_sys_users_list">
    <input />
    <output>
      <integer name="info_size" />
      <for count="info_size" from="0" name="i">
        <integer array_index="i" name="user_id" />
        <string array_index="i" name="login" />
        <ip_address array_index="i" name="ip4_address" />
        <integer array_index="i" name="mask4" />
        <ip_address array_index="i" name="ip6_address" />
        <integer array_index="i" name="mask6" />
      </for>
    </output>
  </function>



*/
