package fn

// rpcf_edit_once_slink_ex
func x2921(c conn, p Dict) Dict {
	c.Call(0x2921)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2921" name="rpcf_edit_once_slink_ex">
    <input>
      <integer default="0" name="slink_id" />
      <integer default="now()" name="discount_date" />
      <double default="1.0" name="cost_coef" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
