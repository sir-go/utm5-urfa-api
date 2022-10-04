package fn

// rpcf_add_sys_user
func x4415(c conn, p Dict) Dict {
	c.Call(0x4415)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4415" name="rpcf_add_sys_user">
      <input>
        <integer name="user_id" />
        <string name="login" />
        <string name="password" />
        <ip_address name="ip4_address" />
        <integer name="mask4" />
        <ip_address name="ip6_address" />
        <integer name="mask6" />
        <integer name="groups_size" />
        <for count="size(group_id)" from="0" name="i">
          <integer array_index="i" name="group_id" />
        </for>
      </input>
      <output>
        <integer name="result" />
      </output>
    </function>


*/
