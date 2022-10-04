package fn

// rpcf_add_ip_slink_ex
func x2928(c conn, p Dict) Dict {
	c.Call(0x2928)

	// fixme: default value parsing error
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2928" name="rpcf_add_ip_slink_ex">
    <input>
      <integer name="user_id" />
      <integer default="basic_account" name="account_id" />
      <integer name="service_id" />
      <integer default="0" name="tariff_link_id" />
      <integer name="discount_period_id" />
      <integer default="now()" name="start_date" />
      <integer default="max_time()" name="expire_date" />
      <integer default="0" name="unabon" />
      <integer default="0" name="unprepay" />
      <integer default="size(ip_address)" name="ip_groups_count" />
      <for count="size(ip_address)" from="0" name="i">
        <integer array_index="i" name="ip_address" />
        <integer array_index="i" name="mask" />
        <string array_index="i" default="" name="mac" />
        <string array_index="i" default="" name="iptraffic_login" />
        <string array_index="i" default="" name="iptraffic_allowed_cid" />
        <string array_index="i" default="" name="iptraffic_password" />
        <integer array_index="i" default="0" name="ip_not_vpn" />
        <integer array_index="i" default="0" name="dont_use_fw" />
        <integer array_index="i" default="0" name="router_id" />
      </for>
      <integer default="size(quota)" name="quotas_count" />
      <for count="size(quota)" from="0" name="i">
        <integer array_index="i" name="tclass_id" />
        <long array_index="i" name="quota" />
      </for>
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
