package fn

// rpcf_edit_hotspot_slink_ex
func x2942(c conn, p Dict) Dict {
	c.Call(0x2942)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2942" name="rpcf_edit_hotspot_slink_ex">
    <input>
      <integer default="0" name="slink_id" />
        <integer default="now()" name="start_date" />
        <integer default="max_time()" name="expire_date" />
        <integer name="policy_id" />
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
