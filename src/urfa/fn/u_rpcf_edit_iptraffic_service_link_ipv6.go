package fn

// rpcf_edit_iptraffic_service_link_ipv6
func x294f(c conn, p Dict) Dict {
	c.Call(0x294f)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x294f" name="rpcf_edit_iptraffic_service_link_ipv6">
    <input>
      <integer default="0" name="slink_id" />
        <integer default="now()" name="start_date" />
        <integer default="max_time()" name="expire_date" />
        <integer name="policy_id" />
        <double name="cost_coef" />
        <integer default="0" name="house_id" />
        <string default="" name="comment" />
        <integer default="size(ip_address)" name="ip_groups_count" />
        <for count="ip_groups_count" from="0" name="i">
          <integer array_index="i" name="id" />
          <ip_address array_index="i" name="ip_address" />
          <integer array_index="i" name="mask" />
          <string array_index="i" default="" name="mac" />
          <string array_index="i" default="" name="iptraffic_login" />
          <string array_index="i" default="" name="iptraffic_allowed_cid" />
          <string array_index="i" default="" name="iptraffic_password" />
          <string array_index="i" default="" name="pool_name" />
          <integer array_index="i" default="0" name="ip_not_vpn" />
          <integer array_index="i" default="0" name="dont_use_fw" />
          <integer array_index="i" default="0" name="router_id" />
          <integer array_index="i" default="0" name="switch_id" />
          <integer array_index="i" default="0" name="port_id" />
          <integer array_index="i" default="0" name="vlan_id" />
          <integer array_index="i" default="0" name="pool_id" />

          <integer array_index="i" default="0" name="dhcp_option_owner_id" />
          <integer array_index="i" default="0" name="dhcp_option_owner_type" />
          <integer array_index="i" default="0" name="dhcp_options_size" />
          <for count="dhcp_options_size" from="0" name="ii">
            <integer array_index="i,ii" default="0" name="dhcp_option_id" />
            <integer array_index="i,ii" default="0" name="dhcp_data_type" />
            <if condition="eq" value="1" variable="dhcp_data_type">
              <integer array_index="i,ii" default="0" name="dhcp_attr_data_int" />
            </if>
            <if condition="eq" value="2" variable="dhcp_data_type">
              <string array_index="i,ii" default="" name="dhcp_attr_data_string" />
            </if>
            <if condition="eq" value="3" variable="dhcp_data_type">
              <ip_address array_index="i,ii" default="0" name="dhcp_attr_data_ip" />
            </if>
            <if condition="eq" value="4" variable="dhcp_data_type">
              <string array_index="i,ii" default="" name="dhcp_attr_data_hex_bin" />
            </if>
            <if condition="eq" value="5" variable="dhcp_data_type">
              <integer array_index="i,ii" default="0" name="ip_array_size" />
              <for count="ip_array_size" from="0" name="iii">
                <ip_address array_index="i,ii,iii" name="attr_data_ip" />
              </for>
            </if>
          </for>
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
