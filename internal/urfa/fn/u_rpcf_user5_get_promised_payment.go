package fn

// rpcf_user5_get_promised_payment
func xU15031(c conn, p Dict) Dict {
	c.Call(-0x15031)
	putI(c, p, "account_id", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15031" name="rpcf_user5_get_promised_payment">
    <input>
      <integer default="0" name="account_id" />
    </input>
    <output>
      <integer name="res" />
        <if condition="ne" value="-1" variable="res">
		<integer name="last_payment_date" />
		<double name="amount" />
		<integer name="duration" />
		<integer name="interval_duration" />
		<double name="cost" />
		<double name="min_balance" />
		<integer name="use_min_balance" />
		<double name="free_balance" />
		<integer name="use_free_balance" />
		<double name="balance" />
                <integer name="flags" />
        </if>
    </output>
  </function>


*/
