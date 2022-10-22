package fn

// rpcf_edit_ip_slink_ex
func x2929(c conn, p Dict) Dict {
	c.Call(0x2929)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2929" name="rpcf_edit_ip_slink_ex">
    <input>
      <integer default="0" name="slink_id" />
      <integer default="now()" name="start_date" />
      <integer default="max_time()" name="expire_date" />

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
