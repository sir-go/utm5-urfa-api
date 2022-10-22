package fn

// rpcf_get_dhcp_leases_report
func x1354(c conn, p Dict) Dict {
	c.Call(0x1354)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "apid")
	putI(c, p, "t_start")
	putI(c, p, "t_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1354" name="rpcf_get_dhcp_leases_report">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="apid" />
      <integer name="t_start" />
      <integer name="t_end" />
    </input>
    <output>
      <integer name="lases_size" />
      <for count="lases_size" from="0" name="i">
          <integer array_index="i" name="user_id" />
          <string array_index="i" name="user_login" />
          <ip_address array_index="i" name="ip" />
          <string array_index="i" name="mac" />
          <string array_index="i" name="relay_agent_info" />
          <integer array_index="i" name="updated" />
          <integer array_index="i" name="expired" />
      </for>
    </output>
  </function>


*/
