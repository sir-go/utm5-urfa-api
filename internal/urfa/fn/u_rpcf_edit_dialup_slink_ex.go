package fn

// rpcf_edit_dialup_slink_ex
func x2945(c conn, p Dict) Dict {
	c.Call(0x2945)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"slink_id": c.GetI(),
	}
}

/* xml:
<function id="0x2945" name="rpcf_edit_dialup_slink_ex">
    <input>
      <integer default="0" name="slink_id" />
        <integer default="now()" name="start_date" />
        <integer default="max_time()" name="expire_date" />
        <integer name="policy_id" />
        <double name="cost_coef" />
        <integer default="0" name="house_id" />
        <string default="" name="comment" />
        <string name="dialup_login" />
        <string default="" name="dialup_password" />
        <string default="" name="dialup_allowed_cid" />
        <string default="" name="dialup_allowed_csid" />
        <integer default="0" name="callback_enabled" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
