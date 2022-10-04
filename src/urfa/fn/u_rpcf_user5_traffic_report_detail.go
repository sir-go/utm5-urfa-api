package fn

// rpcf_user5_traffic_report_detail
func xU4031(c conn, p Dict) Dict {
	c.Call(-0x4031)
	putI(c, p, "limit_size")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4031" name="rpcf_user5_traffic_report_detail">
    <input>
      <integer name="limit_size" />
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="nf5a_size" />
      <for count="nf5a_size" from="0" name="i">
        <integer array_index="i" name="timestamp" />
        <integer array_index="i" name="slink_id" />
        <string array_index="i" name="sname" />
        <integer array_index="i" name="account_id" />
        <integer array_index="i" name="tclass" />
        <string array_index="i" name="tcd_name" />
        <integer array_index="i" name="saddr" />
        <integer array_index="i" name="daddr" />
        <integer array_index="i" name="d_pkt" />
        <integer array_index="i" name="d_okt" />
        <integer array_index="i" name="sport" />
        <integer array_index="i" name="dport" />
        <integer array_index="i" name="tcp_flags" />
        <integer array_index="i" name="proto" />
        <integer array_index="i" name="tos" />
      </for>
    </output>
  </function>


*/
