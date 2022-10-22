package fn

// rpcf_get_userinfo
func x2006(c conn, p Dict) Dict {
	c.Call(0x2006)
	putI(c, p, "user_id")
	c.Send()

	if c.GetI() == 0 {
		return Dict{"error": Dict{
			"code":    10,
			"comment": "user not found"}}
	}

	return Dict{
		"accounts": getMapOf(c, func() Dict {
			return Dict{
				"account_id":   c.GetI(),
				"account_name": c.GetS(),
			}
		}),
		"login":            c.GetS(),
		"password":         c.GetS(),
		"basic_account":    c.GetI(),
		"full_name":        c.GetS(),
		"create_date":      c.GetI(),
		"last_change_date": c.GetI(),
		"who_create":       c.GetI(),
		"who_change":       c.GetI(),
		"is_juridical":     c.GetI(),
		"jur_address":      c.GetS(),
		"act_address":      c.GetS(),
		"work_tel":         c.GetS(),
		"home_tel":         c.GetS(),
		"mob_tel":          c.GetS(),
		"web_page":         c.GetS(),
		"icq_number":       c.GetS(),
		"tax_number":       c.GetS(),
		"kpp_number":       c.GetS(),
		"bank_id":          c.GetI(),
		"bank_account":     c.GetS(),
		"comments":         c.GetS(),
		"personal_manager": c.GetS(),
		"connect_date":     c.GetI(),
		"email":            c.GetS(),
		"is_send_invoice":  c.GetI(),
		"advance_payment":  c.GetI(),
		"house_id":         c.GetI(),
		"flat_number":      c.GetS(),
		"entrance":         c.GetS(),
		"floor":            c.GetS(),
		"district":         c.GetS(),
		"building":         c.GetS(),
		"passport":         c.GetS(),
		"parameters": getMapOf(c, func() Dict {
			return Dict{
				"id":    c.GetI(),
				"value": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2006" name="rpcf_get_userinfo">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="user_id" />
      <if condition="eq" value="0" variable="user_id">
        <error code="10" comment="user not found" />
      </if>
      <integer name="accounts_count" />
      <for count="accounts_count" from="0" name="i">
        <integer array_index="i" name="account_id_array" />
        <string array_index="i" name="account_name_array" />
      </for>
      <string name="login" />
      <string name="password" />
      <integer name="basic_account" />
      <string name="full_name" />
      <integer name="create_date" />
      <integer name="last_change_date" />
      <integer name="who_create" />
      <integer name="who_change" />
      <integer name="is_juridical" />
      <string name="jur_address" />
      <string name="act_address" />
      <string name="work_tel" />
      <string name="home_tel" />
      <string name="mob_tel" />
      <string name="web_page" />
      <string name="icq_number" />
      <string name="tax_number" />
      <string name="kpp_number" />
      <integer name="bank_id" />
      <string name="bank_account" />
      <string name="comments" />
      <string name="personal_manager" />
      <integer name="connect_date" />
      <string name="email" />
      <integer name="is_send_invoice" />
      <integer name="advance_payment" />
      <integer name="house_id" />
      <string name="flat_number" />
      <string name="entrance" />
      <string name="floor" />
      <string name="district" />
      <string name="building" />
      <string name="passport" />
      <integer name="parameters_size" />
      <for count="parameters_size" from="0" name="i">
        <integer array_index="i" name="parameter_id" />
        <string array_index="i" name="parameter_value" />
      </for>
    </output>
  </function>


*/
