package fn

// rpcf_add_tclass2
func x2305(c conn, p Dict) Dict {
	c.Call(0x2305)

	// fixme: input has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2305" name="rpcf_add_tclass2">
    <input>
      <integer name="tclass_id" />
      <string name="tclass_name" />
      <integer name="graph_color" />
      <integer name="is_display" />
      <integer name="is_fill" />
      <integer name="time_range_id" />
      <integer name="dont_save" />
      <integer name="local_traf_policy" />
      <integer name="tclass_count" />
      <for count="tclass_count" from="0" name="i">
        <ip_address array_index="i" name="saddr" />
        <integer array_index="i" name="saddr_mask" />
        <integer array_index="i" name="sport" />
        <integer array_index="i" name="input" />
        <integer array_index="i" name="src_as" />
        <ip_address array_index="i" name="daddr" />
        <integer array_index="i" name="daddr_mask" />
        <integer array_index="i" name="dport" />
        <integer array_index="i" name="output" />
        <integer array_index="i" name="dst_as" />
        <integer array_index="i" name="proto" />
        <integer array_index="i" name="tos" />
        <ip_address array_index="i" name="nexthop" />
        <integer array_index="i" name="tcp_flags" />
        <ip_address array_index="i" name="ip_from" />
        <integer array_index="i" name="nfprovider_id" />
        <integer array_index="i" name="use_sport" />
        <integer array_index="i" name="use_input" />
        <integer array_index="i" name="use_src_as" />
        <integer array_index="i" name="use_dport" />
        <integer array_index="i" name="use_output" />
        <integer array_index="i" name="use_dst_as" />
        <integer array_index="i" name="use_proto" />
        <integer array_index="i" name="use_tos" />
        <integer array_index="i" name="use_nexthop" />
        <integer array_index="i" name="use_tcp_flags" />
        <integer array_index="i" name="skip" />
      </for>
    </input>
    <output />
  </function>


*/
