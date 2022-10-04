package fn

// rpcf_add_user_new
func x2125(c conn, p Dict) Dict {
	c.Call(0x2125)
	putS(c, p, "login")
	putS(c, p, "password")
	putS(c, p, "full_name", "")
	putI(c, p, "is_juridical", 0)
	putS(c, p, "jur_address", "")
	putS(c, p, "act_address", "")
	putS(c, p, "flat_number", "")
	putS(c, p, "entrance", "")
	putS(c, p, "floor", "")
	putS(c, p, "district", "")
	putS(c, p, "building", "")
	putS(c, p, "passport", "")
	putI(c, p, "house_id", 0)
	putS(c, p, "work_tel", "")
	putS(c, p, "home_tel", "")
	putS(c, p, "mob_tel", "")
	putS(c, p, "web_page", "")
	putS(c, p, "icq_number", "")
	putS(c, p, "tax_number", "")
	putS(c, p, "kpp_number", "")
	putS(c, p, "email", "")
	putI(c, p, "bank_id", 0)
	putS(c, p, "bank_account", "")
	putS(c, p, "comments", "")
	putS(c, p, "personal_manager", "")
	putI(c, p, "connect_date", TimeMin)
	putI(c, p, "is_send_invoice", 0)
	putI(c, p, "advance_payment", 0)
	putI(c, p, "switch_id", 0)
	putI(c, p, "port_number", 0)
	putI(c, p, "binded_currency_id", 810)
	forEachDict(c, p, "parameters", func(ai Dict) {
		putI(c, ai, "id")
		putS(c, ai, "value")
	})

	// "groups": [1, 3, ...] - groups id array
	mustHave(p, "groups")
	c.PutInt(len(p["groups"].([]interface{})))
	for _, groupId := range p["groups"].([]interface{}) {
		c.PutInt(int(groupId.(float64)))
	}

	putI(c, p, "is_blocked", 0)
	putD(c, p, "balance", 0.0)
	putD(c, p, "credit", 0.0)
	putD(c, p, "vat_rate", 0.0)
	putD(c, p, "sale_tax_rate", 0.0)
	putI(c, p, "int_status", 1)
	c.Send()

	if uid := c.GetI(); uid != 0 {
		return Dict{"user_id": uid, "basic_account": c.GetI()}
	}
	return Dict{"error_code": c.GetI(), "error_description": c.GetS()}
}

/* xml:
<function id="0x2125" name="rpcf_add_user_new">
    <input>
      <string name="login" />
      <string name="password" />
      <string default="" name="full_name" />
      <integer default="0" name="is_juridical" />
      <string default="" name="jur_address" />
      <string default="" name="act_address" />
      <string default="" name="flat_number" />
      <string default="" name="entrance" />
      <string default="" name="floor" />
      <string default="" name="district" />
      <string default="" name="building" />
      <string default="" name="passport" />
      <integer default="0" name="house_id" />
      <string default="" name="work_tel" />
      <string default="" name="home_tel" />
      <string default="" name="mob_tel" />
      <string default="" name="web_page" />
      <string default="" name="icq_number" />
      <string default="" name="tax_number" />
      <string default="" name="kpp_number" />
      <string default="" name="email" />
      <integer default="0" name="bank_id" />
      <string default="" name="bank_account" />
      <string default="" name="comments" />
      <string default="" name="personal_manager" />
      <integer default="0" name="connect_date" />
      <integer default="0" name="is_send_invoice" />
      <integer default="0" name="advance_payment" />

      <integer default="0" name="switch_id" />
      <integer default="0" name="port_number" />
      <integer default="810" name="binded_currency_id" />

      <integer default="size(parameter_value)" name="parameters_count" />
      <for count="size(parameter_value)" from="0" name="i">
        <integer array_index="i" name="parameter_id" />
        <string array_index="i" name="parameter_value" />
      </for>

      <integer default="size(groups)" name="groups_count" />
      <for count="size(groups)" from="0" name="i">
          <integer array_index="i" name="groups" />
      </for>

      <integer default="0" name="is_blocked" />
      <double default="0.0" name="balance" />
      <double default="0.0" name="credit" />
      <double default="0.0" name="vat_rate" />
      <double default="0.0" name="sale_tax_rate" />
      <integer default="1" name="int_status" />

    </input>
    <output>
      <integer name="user_id" />
      <if condition="eq" value="0" variable="user_id">
          <integer name="error_code" />
          <string name="error_description" />
      </if>
      <if condition="ne" value="0" variable="user_id">
          <integer name="basic_account" />
      </if>
    </output>
  </function>


*/
