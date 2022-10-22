package fn

// rpcf_add_once_slink_ex
func x2920(c conn, p Dict) Dict {
	c.Call(0x2920)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2920" name="rpcf_add_once_slink_ex">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="service_id" />
      <integer default="0" name="tariff_link_id" />
      <integer default="now()" name="discount_date" />
      <double default="1.0" name="cost_coef" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
