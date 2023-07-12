package fn

// rpcf_get_iptraffic_service_link_ipv6
func x272f(c conn, p Dict) Dict {
	c.Call(0x272f)
	putI(c, p, "slink_id")
	c.Send()

	res := Dict{
		"tariff_link_id":     c.GetI(),
		"is_blocked":         c.GetI(),
		"discount_period_id": c.GetI(),
		"start_date":         c.GetI(),
		"expire_date":        c.GetI(),
		"policy_id":          c.GetI(),
		"cost_coef":          c.GetD(),
		"unabon":             c.GetI(),
		"unprepay":           c.GetI(),
		"house_id":           c.GetI(),
		"comment":            c.GetS(),
		"tariff_id":          c.GetI(),
		"parent_id":          c.GetI(),
		"bandwidth_in":       c.GetI(),
		"bandwidth_out":      c.GetI(),
		"ip_groups": getMapOf(c, func() Dict {
			return Dict{
				"id":                    c.GetI(),
				"ip_address":            c.GetA(),
				"mask":                  c.GetI(),
				"mac":                   c.GetS(),
				"iptraffic_login":       c.GetS(),
				"iptraffic_password":    c.GetS(),
				"iptraffic_allowed_cid": c.GetS(),
				"pool_name":             c.GetS(),
				"ip_not_vpn":            c.GetI(),
				"dont_use_fw":           c.GetI(),
				"is_dynamic":            c.GetI(),
				"router_id":             c.GetI(),
				"switch_id":             c.GetI(),
				"port_id":               c.GetI(),
				"vlan_id":               c.GetI(),
				"pool_id":               c.GetI(),
				"dhcp_options": getMapOf(c, func() Dict {
					dhcpOption := Dict{
						"dhcp_option_id": c.GetI(),
						"dhcp_data_type": c.GetI(),
					}
					switch dhcpOption["dhcp_data_type"] {
					case 1:
						dhcpOption["dhcp_attr_data_int"] = c.GetI()
					case 2:
						dhcpOption["dhcp_attr_data_string"] = c.GetS()
					case 3:
						dhcpOption["dhcp_attr_data_ip"] = c.GetA()
					case 4:
						dhcpOption["dhcp_attr_data_hex_bin"] = c.GetBin()
					case 5:
						dhcpOption["ip_array"] = getMapOf(c, func() Dict {
							return Dict{"attr_data_ip": c.GetA()}
						})
					}
					return dhcpOption
				}),
				"quotas": getMapOf(c, func() Dict {
					return Dict{
						"tclass_id":   c.GetI(),
						"tclass_name": c.GetS(),
						"quota":       c.GetL(),
					}
				}),
			}
		}),
	}

	return res
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
