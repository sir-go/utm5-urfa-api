package fn

// rpcf_dealer_add_service_to_user
func x70000030(c conn, p Dict) Dict {
	c.Call(0x70000030)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x70000030" name="rpcf_dealer_add_service_to_user">
      <input>
        <integer name="user_id" />
        <integer name="account_id" />
        <integer name="service_id" />
        <integer name="service_type" />
        <string default="" name="unused" />
        <integer default="0" name="tariff_link_id" />

        <if condition="eq" value="1" variable="service_type">
              <integer default="now()" name="discount_date" />
              <double default="1.0" name="cost_coef" />
        </if>

        <if condition="eq" value="2" variable="service_type">
          <integer name="discount_period_id" />
          <integer default="now()" name="start_date" />
          <integer default="max_time()" name="expire_date" />
          <integer name="policy_id" />
          <integer default="0" name="unabon" />
          <double default="1" name="cost_coef" />
        </if>

        <if condition="eq" value="3" variable="service_type">
          <integer name="discount_period_id" />
          <integer default="now()" name="start_date" />
          <integer default="max_time()" name="expire_date" />
          <integer name="policy_id" />
          <integer default="0" name="unabon" />
          <double default="1" name="cost_coef" />
          <integer default="0" name="unprepay" />
          <integer default="size(ip_address)" name="ip_groups_count" />
          <for count="size(ip_address)" from="0" name="i">
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
          </for>
          <integer default="size(quota)" name="quotas_count" />
          <for count="size(quota)" from="0" name="i">
            <integer array_index="i" name="tclass_id" />
            <long array_index="i" name="quota" />
          </for>
        </if>

        <if condition="eq" value="4" variable="service_type">
          <integer name="discount_period_id" />
          <integer default="now()" name="start_date" />
          <integer default="max_time()" name="expire_date" />
          <integer name="policy_id" />
          <integer default="0" name="unabon" />
          <double default="1" name="cost_coef" />
          <string name="hotspot_login" />
          <string name="hotspot_password" />
          <integer default="0" name="unabon" />
        </if>

        <if condition="eq" value="5" variable="service_type">
          <integer name="discount_period_id" />
          <integer default="now()" name="start_date" />
          <integer default="max_time()" name="expire_date" />
          <integer name="policy_id" />
          <integer default="0" name="unabon" />
          <double default="1" name="cost_coef" />
          <string name="dialup_login" />
          <string default="" name="dialup_password" />
          <string default="" name="dialup_allowed_cid" />
          <string default="" name="dialup_allowed_csid" />
          <integer default="3" name="callback_enabled" />
          <integer default="0" name="unabon" />
        </if>

        <if condition="eq" value="6" variable="service_type">
          <integer name="discount_period_id" />
          <integer default="now()" name="start_date" />
          <integer default="max_time()" name="expire_date" />
          <integer name="policy_id" />
          <integer default="0" name="unabon" />
          <double default="1" name="cost_coef" />
          <integer default="0" name="unabon" />
          <integer default="size(tel_number)" name="tel_numbers_count" />
          <for count="size(tel_number)" from="0" name="i">
            <integer default="0" name="tel_slink_id" />
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
