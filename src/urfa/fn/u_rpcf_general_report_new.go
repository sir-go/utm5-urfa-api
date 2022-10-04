package fn

// rpcf_general_report_new
func x3025(c conn, p Dict) Dict {
	c.Call(0x3025)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3025" name="rpcf_general_report_new">
    <input>
      <integer default="0" name="user_id" />
      <integer default="0" name="account_id" />
      <integer default="0" name="group_id" />
      <integer default="0" name="discount_period_id" />
      <integer name="start_date" />
      <integer default="now()" name="end_date" />
    </input>
    <output>
      <integer name="accounts_count" />
      <for count="accounts_count" from="0" name="i">
        <integer array_index="i" name="account_id" />
        <string array_index="i" name="login" />
        <string array_index="i" name="full_name" />
        <double array_index="i" name="incoming_rest" />
        <double array_index="i" name="discounted_once" />
        <double array_index="i" name="discounted_periodic" />
        <double array_index="i" name="discounted_iptraffic" />
        <double array_index="i" name="discounted_hotspot" />
        <double array_index="i" name="discounted_dialup" />
        <double array_index="i" name="discounted_telephony" />
        <double array_index="i" name="discounted_iptv" />
        <double array_index="i" name="discounted_vod" />
        <double array_index="i" name="discounted_other_charges" />
        <double array_index="i" name="tax" />
        <double array_index="i" name="discounted_with_tax" />
        <double array_index="i" name="payments" />
        <double array_index="i" name="outgoing_rest" />
      </for>
    </output>
  </function>


*/
