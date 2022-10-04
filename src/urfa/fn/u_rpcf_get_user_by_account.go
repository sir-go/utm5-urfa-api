package fn

// rpcf_get_user_by_account
func x2026(c conn, p Dict) Dict {
	c.Call(0x2026)
	putI(c, p, "account_id")
	c.Send()

	if uid := c.GetI(); uid != 0 {
		return Dict{"user_id": uid}
	}

	return Dict{"error": Dict{"code": 19, "comment": "No such account linked with user"}}
}

/* xml:
<function id="0x2026" name="rpcf_get_user_by_account">
    <input>
      <integer name="account_id" />
    </input>
    <output>
      <integer name="user_id" />
      <if condition="eq" value="0" variable="user_id">
        <error code="19" comment="No such account linked with user" />
      </if>
    </output>
  </function>


*/
