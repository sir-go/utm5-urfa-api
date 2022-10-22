package fn

// rpcf_add_tel_service_link
func x1033a(c conn, p Dict) Dict {
	c.Call(0x1033a)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x1033a" name="rpcf_add_tel_service_link">
	  <input>
		  <integer name="user_id" />
		  <integer name="account_id" />
		  <integer name="service_id" />
		  <integer name="tplink_id" />
		  <integer name="discount_period_id" />
		  <integer name="start_date" />
		  <integer name="expire_date" />
          <integer name="policy_id" />
		  <integer name="recalc_fee" />
          <double default="1.0" name="cost_coef" />
          <integer default="0" name="house_id" />
          <string default="" name="comment" />
		  <integer name="login_count" />
		  <for count="login_count" from="0" name="i">
			  <string array_index="i" name="login" />
			  <string array_index="i" name="number" />
			  <string array_index="i" name="incoming_trunk" />
			  <string array_index="i" name="outgoing_trunk" />
			  <string array_index="i" name="pbx_id" />
			  <string array_index="i" name="password" />
			  <string array_index="i" name="allowed_cid" />
		  </for>
	  </input>
	  <output>
		  <integer name="slink_id" />
	  </output>
  </function>


*/
