package fn

// rpcf_add_payment_for_account
func x3110(c conn, p Dict) Dict {
	c.Call(0x3110)
	putI(c, p, "account_id")
	putI(c, p, "unused", 0)
	putD(c, p, "payment")
	putI(c, p, "currency_id", 810)
	putI(c, p, "payment_date", TimeNow())
	putI(c, p, "burn_date", 0)
	putI(c, p, "payment_method", 1)
	putS(c, p, "admin_comment", "")
	putS(c, p, "comment", "")
	putS(c, p, "payment_ext_number")
	putI(c, p, "payment_to_invoice", 0)
	putI(c, p, "turn_on_inet", 1)
	c.Send()

	return Dict{"payment_transaction_id": c.GetI()}
}

/* xml:
<function id="0x3110" name="rpcf_add_payment_for_account">
    <input>
      <integer name="account_id" />
      <integer default="0" name="unused" />
      <double name="payment" />
      <integer default="810" name="currency_id" />
      <integer default="now()" name="payment_date" />
      <integer default="0" name="burn_date" />
      <integer default="1" name="payment_method" />
      <string default="" name="admin_comment" />
      <string default="" name="comment" />
      <string default="" name="payment_ext_number" />
      <integer default="0" name="payment_to_invoice" />
      <integer default="1" name="turn_on_inet" />
    </input>
    <output>
      <integer name="payment_transaction_id" />
    </output>
  </function>


*/
