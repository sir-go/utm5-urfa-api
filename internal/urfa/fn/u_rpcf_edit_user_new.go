package fn

// rpcf_edit_user_new
func x2126(c conn, p Dict) Dict {
	c.Call(0x2126)
	putI(c, p, "user_id")
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
	c.Send()

	if uid := c.GetI(); uid != 0 {
		return Dict{"user_id": uid}
	}
	return Dict{"error_code": c.GetI(), "error_description": c.GetS()}
}

/* xml:
<function id="0x2126" name="rpcf_edit_user_new">
    <input>
      <integer name="user_id" />
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

    </input>
    <output>
      <integer name="user_id" />
      <if condition="eq" value="0" variable="user_id">
          <integer name="error_code" />
	  <string name="error_description" />
      </if>
    </output>
  </function>


*/
