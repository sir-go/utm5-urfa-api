package fn

// rpcf_add_periodic_slink_ex
func x2946(c conn, p Dict) Dict {
	c.Call(0x2946)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "service_id")
	putI(c, p, "tariff_link_id", 0)
	putI(c, p, "discount_period_id")
	putI(c, p, "start_date", TimeNow())
	putI(c, p, "expire_date", TimeMax)
	putI(c, p, "policy_id")
	putI(c, p, "house_id", 0)
	putS(c, p, "comment", "")
	putI(c, p, "unabon", 0)
	putD(c, p, "cost_coef", 1.0)
	c.Send()

	return Dict{"slink_id": c.GetI()}
}

/* xml:
<function id="0x2946" name="rpcf_add_periodic_slink_ex">
    <input>
		<integer name="user_id" />
		<integer name="account_id" />
		<integer name="service_id" />
		<integer default="0" name="tariff_link_id" />
		<integer name="discount_period_id" />
		<integer default="now()" name="start_date" />
		<integer default="max_time()" name="expire_date" />
		<integer name="policy_id" />
		<integer default="0" name="house_id" />
		<string default="" name="comment" />
		<integer default="0" name="unabon" />
		<double default="1" name="cost_coef" />
    </input>
    <output>
      <integer name="slink_id" />
    </output>
  </function>


*/
