package fn

// rpcf_get_traffic_detailed
func x502b(c conn, p Dict) Dict {
	c.Call(0x502b)
	putI(c, p, "stop")
	putI(c, p, "apid")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x502b" name="rpcf_get_traffic_detailed">
    <input>
      <integer name="stop" />
      <integer name="apid" />
    </input>
    <output>
      <integer name="nf5a_size" />
      <integer name="is_finish" />
      <for count="nf5a_size" from="0" name="i">
        <integer array_index="i" name="timestamp" />
        <integer array_index="i" name="slink_id" />
        <integer array_index="i" name="account_id" />
        <integer array_index="i" name="tclass" />
        <ip_address array_index="i" name="saddr" />
        <ip_address array_index="i" name="daddr" />
        <integer array_index="i" name="d_pkt" />
        <integer array_index="i" name="d_oct" />
        <integer array_index="i" name="sport" />
        <integer array_index="i" name="dport" />
        <integer array_index="i" name="tcp_flags" />
        <integer array_index="i" name="proto" />
        <integer array_index="i" name="tos" />
      </for>
    </output>
  </function>


*/
