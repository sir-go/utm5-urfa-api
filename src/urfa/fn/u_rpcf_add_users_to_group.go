package fn

// rpcf_add_users_to_group
func x2407(c conn, p Dict) Dict {
	c.Call(0x2407)
	putI(c, p, "group_id")

	mustHave(p, "user_ids")
	forEach(c, p, "user_ids", func(ai interface{}) {
		c.PutInt(int(ai.(float64)))
	})
	c.Send()
	return nil
}

/* xml:
<function id="0x2407" name="rpcf_add_users_to_group">
    <input>
      <integer name="group_id" />
      <integer default="size(user_id)" name="count" />
      <for count="size(user_id)" from="0" name="i">
        <integer array_index="i" name="user_id" />
      </for>
    </input>
    <output />
  </function>


*/
