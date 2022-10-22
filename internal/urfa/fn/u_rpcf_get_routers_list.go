package fn

// rpcf_get_routers_list
func x5043(c conn, _ Dict) Dict {
	c.Call(0x5043)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5043" name="rpcf_get_routers_list">
    <input />
    <output>
      <integer name="routers_size" />
      <for count="routers_size" from="0" name="i">
        <integer array_index="i" name="router_id" />
        <integer array_index="i" name="router_type" />
        <string array_index="i" name="router_ip" />
        <string array_index="i" name="login" />
        <string array_index="i" name="password" />
        <string array_index="i" name="router_comments" />
        <ip_address array_index="i" name="router_bin_ip" />
        <integer array_index="i" name="is_online" />
      </for>
    </output>
  </function>


*/
