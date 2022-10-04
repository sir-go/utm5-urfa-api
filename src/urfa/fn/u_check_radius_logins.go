package fn

// check_radius_logins
func x293c(c conn, p Dict) Dict {
	c.Call(0x293c)
	putI(c, p, "slink_id")
	forEach(c, p, "logins", func(item interface{}) {
		c.PutStr(item.(string))
	})
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x293c" name="check_radius_logins">
      <input>
        <integer name="slink_id" />
        <integer name="size" />
        <for count="size" from="0" name="i">
          <string array_index="i" name="login" />
        </for>
      </input>
      <output>
        <integer name="result" />
      </output>
    </function>


*/
