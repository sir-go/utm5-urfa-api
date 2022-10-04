package fn

// rpcf_add_iptraffic_service_link_ipv6
func x294e(c conn, p Dict) Dict {
	c.Call(0x294e)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x294e" name="rpcf_add_iptraffic_service_link_ipv6">
        <input>
            <integer name="user_id" />
            <integer name="account_id" />
            <integer name="service_id" />
            <integer name="tplink_id" />
            <integer name="discount_period_id" />
            <integer default="now()" name="start_date" />
            <integer default="max_time()" name="expire_date" />
            <integer name="policy_id" />
            <integer name="unabon" />
            <double default="1.0" name="cost_coef" />
            <integer default="0" name="house_id" />
            <string default="" name="comment" />
            <integer name="unprepay" />
            <integer name="ip_groups_count" />
            <for count="ip_groups_count" from="0" name="i">
                <integer array_index="i" name="id" />
                <ip_address array_index="i" name="ip" />
                <integer array_index="i" name="mask" />
                <string array_index="i" name="mac" />
                <string array_index="i" name="login" />
                <string array_index="i" name="allowed_cid" />
                <string array_index="i" name="password" />
                <string array_index="i" default="" name="pool_name" />
                <integer array_index="i" name="is_skip_radius" />
                <integer array_index="i" name="is_skip_rfw" />
                <integer array_index="i" name="router_id" />
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
                <integer array_index="i" name="tc_id" />
                <long array_index="i" name="quota" />
            </for>
        </input>
        <output>
            <integer name="slink_id" />
        </output>
    </function>


*/
