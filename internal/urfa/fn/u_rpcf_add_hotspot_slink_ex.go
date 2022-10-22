package fn

// rpcf_add_hotspot_slink_ex
func x2943(c conn, p Dict) Dict {
	c.Call(0x2943)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2943" name="rpcf_add_hotspot_slink_ex">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="service_id" />
      <integer default="0" name="tariff_link_id" />
      <integer name="discount_period_id" />
      <integer default="now()" name="start_date" />
      <integer default="max_time()" name="expire_date" />
      <integer name="policy_id" />
      <integer default="0" name="unabon" />
      <double default="1.0" name="cost_coef" />
      <integer default="0" name="house_id" />
      <string default="" name="comment" />
      <string name="hotspot_login" />
      <string name="hotspot_password" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
