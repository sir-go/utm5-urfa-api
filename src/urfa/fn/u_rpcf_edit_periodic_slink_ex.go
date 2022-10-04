package fn

// rpcf_edit_periodic_slink_ex
func x2947(c conn, p Dict) Dict {
	c.Call(0x2947)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2947" name="rpcf_edit_periodic_slink_ex">
    <input>
      <integer default="0" name="slink_id" />
        <integer default="now()" name="start_date" />
        <integer default="max_time()" name="expire_date" />
        <integer name="policy_id" />
        <double name="cost_coef" />
        <integer default="0" name="house_id" />
        <string default="" name="comment" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
