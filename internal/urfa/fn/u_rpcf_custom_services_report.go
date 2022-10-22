package fn

// rpcf_custom_services_report
func x3114(c conn, p Dict) Dict {
	c.Call(0x3114)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3114" name="rpcf_custom_services_report">
    <input>
        <integer default="0" name="t_start" />
        <integer default="now()" name="t_end" />
        <integer default="0" name="account_id" />
        <integer default="0" name="user_id" />
    </input>
    <output>
        <integer name="count" />
        <for count="count" from="0" name="i">
            <integer name="account_id" />
            <string name="login" />
            <integer name="date" />
            <string name="mark" />
            <double name="amount" />
            <double name="amount_with_tax" />
            <string name="service_name" />
            <string name="service_key" />
            <integer name="revoked" />
        </for>
    </output>
  </function>


*/
