package fn

// rpcf_get_iptraffic_service_link_ipv6
func x272f(c conn, p Dict) Dict {
	c.Call(0x272f)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x272f" name="rpcf_get_iptraffic_service_link_ipv6">
    <input>
      <integer name="slink_id" />
    </input>
    <output>
      <integer name="tariff_link_id" />
      <integer name="is_blocked" />
      <integer name="discount_period_id" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <integer name="policy_id" />
      <double name="cost_coef" />
      <integer name="unabon" />
      <integer name="unprepay" />
      <integer name="house_id" />
      <string name="comment" />
      <integer name="tariff_id" />
      <integer name="parent_id" />
      <integer name="bandwidth_in" />
      <integer name="bandwidth_out" />
      <integer name="ip_groups_count" />
      <for count="ip_groups_count" from="0" name="i">
        <integer array_index="i" name="id" />
        <ip_address array_index="i" name="ip_address" />
        <integer array_index="i" name="mask" />
        <string array_index="i" name="mac" />
        <string array_index="i" name="iptraffic_login" />
        <string array_index="i" name="iptraffic_password" />
        <string array_index="i" name="iptraffic_allowed_cid" />
        <string array_index="i" name="pool_name" />
        <integer array_index="i" name="ip_not_vpn" />
        <integer array_index="i" name="dont_use_fw" />
        <integer array_index="i" name="is_dynamic" />
        <integer array_index="i" name="router_id" />
        <integer array_index="i" name="switch_id" />
        <integer array_index="i" name="port_id" />
        <integer array_index="i" name="vlan_id" />
        <integer array_index="i" name="pool_id" />
        <integer array_index="i" default="0" name="dhcp_options_size" />
        <for count="dhcp_options_size" from="0" name="ii">
            <integer array_index="i,ii" default="0" name="dhcp_option_id" />
            <integer array_index="i,ii" default="0" name="dhcp_data_type" />
            <if condition="eq" value="1" variable="dhcp_data_type">
                <integer array_index="i,ii" default="0" name="dhcp_attr_data_int" />
            </if>
            <if condition="eq" value="2" variable="dhcp_data_type">
                <string array_index="i,ii" default="0" name="dhcp_attr_data_string" />
            </if>
            <if condition="eq" value="3" variable="dhcp_data_type">
                <ip_address array_index="i,ii" default="0" name="dhcp_attr_data_ip" />
            </if>
            <if condition="eq" value="4" variable="dhcp_data_type">
                <string array_index="i,ii" default="0" name="dhcp_attr_data_hex_bin" />
            </if>
            <if condition="eq" value="5" variable="dhcp_data_type">
                <integer array_index="i,ii" default="0" name="ip_array_size" />
                <for count="ip_array_size" from="0" name="iii">
                    <ip_address array_index="i,ii,iii" name="attr_data_ip" />
                </for>
            </if>
        </for>
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
