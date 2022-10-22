package fn

// rpcf_user5_get_voluntary_blocking
func xU15014(c conn, p Dict) Dict {
	c.Call(-0x15014)
	putI(c, p, "account_id", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15014" name="rpcf_user5_get_voluntary_blocking">
    <input>
      <integer default="0" name="account_id" />
    </input>
    <output>
      <integer name="res" />
        <if condition="eq" value="1" variable="res">
		<integer name="block_start" />
		<integer name="block_end" />
		<integer name="block_self_unlock" />

        </if>
        <if condition="eq" value="0" variable="res">
		<integer name="can_set" />
		<integer name="last_block_start" />
		<integer name="min_duration" />
		<integer name="max_duration" />
		<integer name="interval_duration" />
		<integer name="block_type" />
		<double name="min_balance" />
		<integer name="use_min_balance" />
		<double name="free_balance" />
		<integer name="use_free_balance" />
		<integer name="self_unlock" />
		<double name="cost" />
		<double name="balance" />

        </if>
    </output>
  </function>


*/
