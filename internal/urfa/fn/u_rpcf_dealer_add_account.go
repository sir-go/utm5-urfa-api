package fn

// rpcf_dealer_add_account
func x70000011(c conn, p Dict) Dict {
	c.Call(0x70000011)
	putI(c, p, "user_id")
	putI(c, p, "is_basic", 1)
	putI(c, p, "is_blocked", 0)
	putS(c, p, "account_name", "auto create account")
	putD(c, p, "balance", 0.0)
	putD(c, p, "credit", 0.0)
	putI(c, p, "discount_period_id", 0)
	putI(c, p, "dealer_account_id", 0)
	putD(c, p, "comission_coefficient", 0.0)
	putD(c, p, "default_comission_value", 0.0)
	putI(c, p, "is_dealer", 0)
	putD(c, p, "vat_rate", 0.0)
	putD(c, p, "sale_tax_rate", 0.0)
	putI(c, p, "int_status", 1)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000011" name="rpcf_dealer_add_account">
      <input>
        <integer name="user_id" />
        <integer default="1" name="is_basic" />
        <integer default="0" name="is_blocked" />
        <string default="auto create account" name="account_name" />
        <double default="0.0" name="balance" />
        <double default="0.0" name="credit" />
        <integer default="0" name="discount_period_id" />
        <integer default="0" name="dealer_account_id" />
        <double default="0.0" name="comission_coefficient" />
        <double default="0.0" name="default_comission_value" />
        <integer default="0" name="is_dealer" />
        <double default="0.0" name="vat_rate" />
        <double default="0.0" name="sale_tax_rate" />
        <integer default="1" name="int_status" />
      </input>
      <output>
        <integer name="account_id" />
        <if condition="eq" value="0" variable="account_id">
          <error code="11" comment="unable to add account or you dont have enough privileges" />
        </if>
      </output>
    </function>


*/
