package fn

// rpcf_link_user_tariff
func x301f(c conn, p Dict) Dict {
	c.Call(0x301f)
	putI(c, p, "user_id")
	putI(c, p, "account_id", 0)
	putI(c, p, "tariff_current")
	putI(c, p, "tariff_next", int(p["tariff_current"].(float64)))
	putI(c, p, "discount_period_id")
	putI(c, p, "tariff_link_id", 0)
	putI(c, p, "change_now", 0)
	c.Send()

	if tlinkId := c.GetI(); tlinkId != 0 {
		return Dict{"tariff_link_id": tlinkId}
	}

	return Dict{"error": Dict{"code": 13, "comment": "unable to link user tariff"}}
}

/* xml:
<function id="0x301f" name="rpcf_link_user_tariff">
    <input>
      <integer name="user_id" />
      <integer default="0" name="account_id" />
      <integer name="tariff_current" />
      <integer default="tariff_current" name="tariff_next" />
      <integer name="discount_period_id" />
      <integer default="0" name="tariff_link_id" />
      <integer default="0" name="change_now" />
    </input>
    <output>
      <integer name="tariff_link_id" />
      <if condition="eq" value="0" variable="tariff_link_id">
        <error code="13" comment="unable to link user tariff" />
      </if>
    </output>
  </function>


*/
