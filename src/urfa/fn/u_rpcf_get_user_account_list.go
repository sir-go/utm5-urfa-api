package fn

// rpcf_get_user_account_list
func x2033(c conn, p Dict) Dict {
	c.Call(0x2033)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"accounts": getMapOf(c, func() Dict {
			return Dict{
				"account":      c.GetI(),
				"account_name": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2033" name="rpcf_get_user_account_list">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="accounts_count" />
    <for count="accounts_count" from="0" name="i">
      <integer array_index="i" name="account" />
      <string array_index="i" name="account_name" />
    </for>
    </output>
  </function>


*/
