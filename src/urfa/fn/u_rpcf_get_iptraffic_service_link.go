package fn

// rpcf_get_iptraffic_service_link
func x2702(c conn, p Dict) Dict {
	c.Call(0x2702)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2702" name="rpcf_get_iptraffic_service_link">
    <input>
      <integer name="slink_id" />
    </input>
    <output>
      <integer name="tariff_link_id" />
      <integer name="is_blocked" />
      <integer name="discount_period_id" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <integer name="unabon" />
      <integer name="unprepay" />
      <integer name="tariff_id" />
      <integer name="parent_id" />
      <integer name="ip_groups_count" />
      <for count="ip_groups_count" from="0" name="i">
        <integer array_index="i" name="ip_address" />
        <integer array_index="i" name="mask" />
        <string array_index="i" name="mac" />
        <string array_index="i" name="iptraffic_login" />
        <string array_index="i" name="iptraffic_password" />
        <string array_index="i" name="iptraffic_allowed_cid" />
        <integer array_index="i" name="ip_not_vpn" />
        <integer array_index="i" name="dont_use_fw" />
        <integer array_index="i" name="router_id" />
      </for>
      <integer name="quotas_count" />
      <for count="quotas_count" from="0" name="i">
        <integer array_index="i" name="tclass_id" />
        <string array_index="i" name="tclass_name" />
        <long array_index="i" name="quota" />
      </for>
    </output>
  </function>


*/
