package fn

// rpcf_dealer_get_once_service_link
func x7000001d(c conn, p Dict) Dict {
	c.Call(0x7000001d)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000001d" name="rpcf_dealer_get_once_service_link">
      <input>
	<integer name="slink_id" />
      </input>
      <output>
         <integer name="sid" />
          <if condition="eq" value="0" variable="sid">
             <error code="11" comment="service link not found or you dont have enough privileges" />
          </if>
	<integer name="discount_date" />
	<double name="quantity" />
	<integer name="invoice_id" />
    <double default="1.0" name="cost_coef" />
      </output>
    </function>


*/
