package fn

// make_payment_for_supplier
func x8018(c conn, p Dict) Dict {
	c.Call(0x8018)
	putI(c, p, "supplier_id")
	putD(c, p, "sum")
	putS(c, p, "reason_of_payment")
	putI(c, p, "currency_type")
	c.Send()

	if supplierId := c.GetI(); supplierId != -1 {
		return Dict{"supplier_id": supplierId, "abs_sum": c.GetD()}
	}
	return Dict{"supplier_id": -1, "err_string": c.GetS()}
}

/* xml:
<function id="0x8018" name="make_payment_for_supplier">
    <input>
      <integer name="supplier_id" />
      <double name="sum" />
      <string name="reason_of_payment" />
      <integer name="currency_type" />
    </input>
    <output>
      <integer name="supplier_id" />
      <if condition="eq" value="-1" variable="supplier_id">
        <string name="err_string" />
      </if>
      <if condition="ne" value="-1" variable="supplier_id">
        <double name="abs_sum" />
      </if>
    </output>
  </function>


*/
