package fn

// rpcf_get_sys_user
func x4414(c conn, p Dict) Dict {
	c.Call(0x4414)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4414" name="rpcf_get_sys_user">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <string name="login" />
      <ip_address name="ip4" />
      <integer name="mask4" />
      <ip_address name="ip6" />
      <integer name="mask6" />
      <integer name="groups_size" />
      <for count="groups_size" from="0" name="i">
        <integer array_index="i" name="group_id" />
        <string array_index="i" name="group_name" />
        <string array_index="i" name="group_info" />
      </for>
    </output>
  </function>


*/
