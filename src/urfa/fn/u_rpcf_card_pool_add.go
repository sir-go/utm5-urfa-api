package fn

// rpcf_card_pool_add
func x4203(c conn, p Dict) Dict {
	c.Call(0x4203)

	// fixme: function in the default value
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4203" name="rpcf_card_pool_add">
    <input>
      <integer default="0" name="pool_id" />
      <integer default="8" name="sec_size" />
      <integer default="0" name="delay" />
      <integer name="size" />
      <double name="balance" />
      <integer default="810" name="currency" />
      <integer default="max_time()" name="expire" />
      <integer default="0" name="days" />
      <integer default="0" name="service_id" />
      <integer default="0" name="random" />
    </input>
    <output />
  </function>


*/
