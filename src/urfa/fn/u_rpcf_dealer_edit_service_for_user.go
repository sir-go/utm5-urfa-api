package fn

// rpcf_dealer_edit_service_for_user
func x7000006c(c conn, p Dict) Dict {
	c.Call(0x7000006c)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x7000006c" name="rpcf_dealer_edit_service_for_user">
      <input>
        <integer name="user_id" />
        <integer default="0" name="account_id" />
        <integer name="service_id" />
        <integer name="service_type" />
        <integer name="tariff_link_id" />

        <if condition="eq" value="1" variable="service_type">
          <integer name="slink_id" />
          <integer name="discount_date" />
          <double default="1.0" name="cost_coef" />
        </if>

        <if condition="eq" value="2" variable="service_type">
          <integer name="slink_id" />
          <integer name="is_blocked" />
          <integer name="discount_period_id" />
          <integer name="start_date" />
          <integer name="expire_date" />
          <integer name="policy_id" />
          <double default="1" name="cost_coef" />
        </if>

        <if condition="eq" value="3" variable="service_type">
          <integer name="slink_id" />
          <integer name="is_blocked" />
          <integer name="discount_period_id" />
          <integer name="start_date" />
          <integer name="expire_date" />
          <integer name="policy_id" />
          <double default="1" name="cost_coef" />
          <integer default="size(ip_address)" name="ip_groups_count" />
          <for count="size(ip_address)" from="0" name="i">
            <ip_address array_index="i" name="ip_address" />
            <integer array_index="i" name="mask" />
            <string array_index="i" name="mac" />
            <string array_index="i" name="iptraffic_login" />
            <string array_index="i" name="iptraffic_allowed_cid" />
            <string array_index="i" name="iptraffic_password" />
            <string array_index="i" name="pool_name" />
            <integer array_index="i" name="ip_not_vpn" />
            <integer array_index="i" name="dont_use_fw" />
            <integer array_index="i" name="router_id" />
            <integer array_index="i" name="switch_id" />
            <integer array_index="i" name="port_id" />
            <integer array_index="i" name="vlan_id" />
            <integer array_index="i" name="pool_id" />
          </for>
          <integer default="size(quota)" name="quotas_count" />
          <for count="size(quota)" from="0" name="i">
            <integer array_index="i" name="tclass_id" />
            <long array_index="i" name="quota" />
          </for>
        </if>

        <if condition="eq" value="4" variable="service_type">
          <integer name="slink_id" />
          <integer name="is_blocked" />
          <integer name="discount_period_id" />
          <integer name="start_date" />
          <integer name="expire_date" />
          <integer name="policy_id" />
          <double default="1" name="cost_coef" />
          <string name="hotspot_login" />
          <string name="hotspot_password" />
        </if>

        <if condition="eq" value="5" variable="service_type">
          <integer name="slink_id" />
          <integer name="is_blocked" />
          <integer name="discount_period_id" />
          <integer name="start_date" />
          <integer name="expire_date" />
          <integer name="policy_id" />
          <double default="1" name="cost_coef" />
          <string name="dialup_login" />
          <string name="dialup_password" />
          <string name="dialup_allowed_cid" />
          <string name="dialup_allowed_csid" />
          <integer name="callback_enabled" />
        </if>

        <if condition="eq" value="6" variable="service_type">
          <integer name="slink_id" />
          <integer name="is_blocked" />
          <integer name="discount_period_id" />
          <integer name="start_date" />
          <integer name="expire_date" />
          <integer name="policy_id" />
          <double default="1" name="cost_coef" />
          <integer default="size(tel_number)" name="tel_numbers_count" />
          <for count="size(tel_number)" from="0" name="i">
	    <integer array_index="i" default="0" name="tel_slink_id" />
            <string array_index="i" name="tel_number" />
            <string array_index="i" name="tel_login" />
            <string array_index="i" name="tel_password" />
            <string array_index="i" default="" name="tel_allowed_cid" />
          </for>
        </if>
      </input>
      <output>
        <integer name="slink_id" />
      </output>
    </function>


*/
