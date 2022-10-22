package fn

// rpcf_search_users_new
func x1206(c conn, p Dict) Dict {
	c.Call(0x1206)

	// "fields": [8, 9, 33 ...]   UserField* constants
	var fields []interface{}

	if _, ok := p["fields"]; ok {
		fields = p["fields"].([]interface{})
	} else {
		fields = make([]interface{}, 0)
	}

	c.PutInt(len(fields))
	for _, fieldCode := range fields {
		c.PutInt(int(fieldCode.(float64)))
	}

	putI(c, p, "select_type", SelectTypeAnd)
	forEachDict(c, p, "patterns", func(bi Dict) {
		putI(c, bi, "what")
		putI(c, bi, "criteria_id", SearchCriteriaEq)

		switch bi["what"] {
		case UserFieldConnectDate:
			fallthrough
		case UserFieldCreateDate:
			fallthrough
		case UserFieldLastChangeDate:
			putI(c, bi, "pattern")
		default:
			putS(c, bi, "pattern")
		}
	})
	c.Send()

	res := Dict{
		"users": getMapOf(c, func() Dict {
			user := Dict{
				"user_id":       c.GetI(),
				"login":         c.GetS(),
				"basic_account": c.GetI(),
				"full_name":     c.GetS(),
				"is_blocked":    c.GetI(),
				"balance":       c.GetD(),
				"ip_addresses": getMapOf(c, func() Dict {
					return Dict{
						"ip_groups": getMapOf(c, func() Dict {
							return Dict{
								"type": c.GetI(),
								"ip":   c.GetA(),
								"mask": c.GetI(),
							}
						}),
					}

				}),
			}

			for _, fieldCode := range fields {
				switch int(fieldCode.(float64)) {

				case UserFieldDiscountPeriodId:
					user["discount_period_id"] = c.GetI()

				case UserFieldCreateDate:
					user["create_date"] = c.GetI()

				case UserFieldLastChangeDate:
					user["last_change_date"] = c.GetI()

				case UserFieldWhoCreate:
					user["who_create"] = c.GetI()

				case UserFieldWhoChange:
					user["who_change"] = c.GetI()

				case UserFieldIsJur:
					user["is_juridical"] = c.GetI()

				case UserFieldJurAddress:
					user["juridical_address"] = c.GetS()

				case UserFieldActualAddress:
					user["actual_address"] = c.GetS()

				case UserFieldWorkPhone:
					user["work_telephone"] = c.GetS()

				case UserFieldHomePhone:
					user["home_telephone"] = c.GetS()

				case UserFieldMobilePhone:
					user["mobile_telephone"] = c.GetS()

				case UserFieldWebPage:
					user["web_page"] = c.GetS()

				case UserFieldIcqUin:
					user["icq_number"] = c.GetS()

				case UserFieldTaxNumber:
					user["tax_number"] = c.GetS()

				case UserFieldKppNumber:
					user["kpp_number"] = c.GetS()

				case UserFieldHouseId:
					user["house_id"] = c.GetI()

				case UserFieldFlatNumber:
					user["flat_number"] = c.GetS()

				case UserFieldEntrance:
					user["entrance"] = c.GetS()

				case UserFieldFloor:
					user["floor"] = c.GetS()

				case UserFieldEmail:
					user["email"] = c.GetS()

				case UserFieldPassport:
					user["passport"] = c.GetS()

				case UserFieldDistrict:
					user["district"] = c.GetS()

				case UserFieldBuilding:
					user["building"] = c.GetS()

				case UserFieldExternalId:
					user["external_id"] = c.GetS()
				}
			}

			return user
		}),
	}

	return res
}

/* xml:
<function id="0x1206" name="rpcf_search_users_new">
    <input>
      <integer default="size(pole_code_array)" name="poles_count" />
      <for count="size(pole_code_array)" from="0" name="i">
        <integer array_index="i" name="pole_code_array" />
      </for>
      <integer name="select_type" />
      <integer default="size(what_id)" name="patterns_count" />
      <for count="size(what_id)" from="0" name="i">
        <set dst="what" src="what_id" src_index="i" />
        <integer name="what" />
        <integer array_index="i" name="criteria_id" />
        <if comment="date" condition="eq" value="33" variable="what">
            <integer array_index="i" name="pattern" />
        </if>
        <if comment="date" condition="eq" value="6" variable="what">
            <integer array_index="i" name="pattern" />
        </if>
        <if comment="date" condition="eq" value="7" variable="what">
            <integer array_index="i" name="pattern" />
        </if>
        <if condition="ne" value="33" variable="what">
            <if condition="ne" value="6" variable="what">
                <if condition="ne" value="7" variable="what">
                    <string array_index="i" name="pattern" />
                </if>
            </if>
        </if>
      </for>
    </input>
    <output>
      <integer name="user_data_size" />
      <for count="user_data_size" from="0" name="i">
        <integer array_index="i" name="user_id" />
        <string array_index="i" name="login" />
        <integer array_index="i" name="basic_account" />
        <string array_index="i" name="full_name" />
        <integer array_index="i" name="is_blocked" />
        <double array_index="i" name="balance" />
        <integer name="ip_address_size" />
        <set dst="ip_address_size_array" dst_index="i" src="ip_address_size" />
        <for count="ip_address_size" from="0" name="j">
          <integer name="ip_group_size" />
          <set dst="ip_group_size_array" dst_index="i,j" src="ip_group_size" />
          <for count="ip_group_size" from="0" name="x">
            <integer array_index="i,j,x" name="type" />
            <ip_address array_index="i,j,x" name="ip" />
            <integer array_index="i,j,x" name="mask" />
          </for>
        </for>
        <for count="size(pole_code_array)" from="0" name="z">
          <set dst="pole_code" src="pole_code_array" src_index="z" />
          <if condition="eq" value="4" variable="pole_code">
            <integer array_index="i" name="discount_period_id" />
          </if>
          <if condition="eq" value="6" variable="pole_code">
            <integer array_index="i" name="create_date" />
          </if>
          <if condition="eq" value="7" variable="pole_code">
            <integer array_index="i" name="last_change_date" />
          </if>
          <if condition="eq" value="8" variable="pole_code">
            <integer array_index="i" name="who_create" />
          </if>
          <if condition="eq" value="9" variable="pole_code">
            <integer array_index="i" name="who_change" />
          </if>
          <if condition="eq" value="10" variable="pole_code">
            <integer array_index="i" name="is_juridical" />
          </if>
          <if condition="eq" value="11" variable="pole_code">
            <string array_index="i" name="juridical_address" />
          </if>
          <if condition="eq" value="12" variable="pole_code">
            <string array_index="i" name="actual_address" />
          </if>
          <if condition="eq" value="13" variable="pole_code">
            <string array_index="i" name="work_telephone" />
          </if>
          <if condition="eq" value="14" variable="pole_code">
            <string array_index="i" name="home_telephone" />
          </if>
          <if condition="eq" value="15" variable="pole_code">
            <string array_index="i" name="mobile_telephone" />
          </if>
          <if condition="eq" value="16" variable="pole_code">
            <string array_index="i" name="web_page" />
          </if>
          <if condition="eq" value="17" variable="pole_code">
            <string array_index="i" name="icq_number" />
          </if>
          <if condition="eq" value="18" variable="pole_code">
            <string array_index="i" name="tax_number" />
          </if>
          <if condition="eq" value="19" variable="pole_code">
            <string array_index="i" name="kpp_number" />
          </if>
          <if condition="eq" value="21" variable="pole_code">
            <integer array_index="i" name="house_id" />
          </if>
          <if condition="eq" value="22" variable="pole_code">
            <string array_index="i" name="flat_number" />
          </if>
          <if condition="eq" value="23" variable="pole_code">
            <string array_index="i" name="entrance" />
          </if>
          <if condition="eq" value="24" variable="pole_code">
            <string array_index="i" name="floor" />
          </if>
          <if condition="eq" value="25" variable="pole_code">
            <string array_index="i" name="email" />
          </if>
          <if condition="eq" value="26" variable="pole_code">
            <string array_index="i" name="passport" />
          </if>
          <if condition="eq" value="40" variable="pole_code">
            <string array_index="i" name="district" />
          </if>
          <if condition="eq" value="41" variable="pole_code">
            <string array_index="i" name="building" />
          </if>
          <if condition="eq" value="44" variable="pole_code">
            <string array_index="i" name="external_id" />
          </if>
        </for>
      </for>
    </output>
  </function>


*/
