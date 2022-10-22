package fn

// rpcf_user5_card_payment_new
func xU4045(c conn, p Dict) Dict {
	c.Call(-0x4045)
	putI(c, p, "account_id")
	putI(c, p, "card_id")
	putS(c, p, "secret")
	c.Send()

	if res := c.GetI(); res != 0 {
		return Dict{"result": res}
	}
	return Dict{"error": c.GetS()}
}

/* xml:
<function id="-0x4045" name="rpcf_user5_card_payment_new">
    <input>
      <integer name="account_id" />
      <integer name="card_id" />
      <string name="secret" />
    </input>
    <output>
      <integer name="result" />
      <if condition="eq" value="0" variable="result">
          <string name="error" />
      </if>
    </output>
  </function>


*/
