package fn

// rpcf_user5_get_services_info
func xU404a(c conn, p Dict) Dict {
	c.Call(-0x404a)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x404a" name="rpcf_user5_get_services_info">
    <input>
      <integer name="slink_id" />
    </input>
    <output>
      <integer name="service_type" />
      <integer name="service_id" />
      <string name="service_name" />
      <integer name="tariff_id" />
      <double name="discounted_in_curr_period" />
      <double name="cost" />
      <if condition="eq" value="3" variable="service_type">
        <integer name="bytes_in_mbyte" />
        <integer name="iptsl_downloaded_size" />
        <for count="iptsl_downloaded_size" from="0" name="i">
          <string array_index="i" name="tclass" />
          <long array_index="i" name="downloaded" />
        </for>
        <integer name="iptsl_old_prepaid_size" />
        <for count="iptsl_old_prepaid_size" from="0" name="i">
          <string array_index="i" name="tclass" />
          <long array_index="i" name="prepaid" />
        </for>
        <integer name="ipgroup_size" />
        <for count="ipgroup_size" from="0" name="i">
          <integer array_index="i" name="item_id" />
          <ip_address array_index="i" name="ip" />
          <integer array_index="i" name="mask" />
          <string array_index="i" name="login" />
        </for>
        <integer name="iptsd_borders_size" />
        <for count="iptsd_borders_size" from="0" name="i">
          <string array_index="i" name="tclass_name" />
          <long array_index="i" name="bytes" />
          <double array_index="i" name="cost1" />
          <integer array_index="i" name="group_type" />
        </for>
        <integer name="iptsd_prepaid_size" />
        <for count="iptsd_prepaid_size" from="0" name="i">
          <string array_index="i" name="tclass_name_p" />
          <long array_index="i" name="prepaid_p" />
        </for>
      </if>
      <if condition="eq" value="6" variable="service_type">
        <integer name="tsl_numbers_size" />
        <for count="tsl_numbers_size" from="0" name="i">
          <string array_index="i" name="number" />
          <string array_index="i" name="login" />
          <string array_index="i" name="allowed_cid" />
          <integer array_index="i" name="item_id" />
        </for>
      </if>
      <if condition="ne" value="3" variable="service_type">
        <if condition="ne" value="6" variable="service_type">
          <integer name="null_param" />
        </if>
      </if>
    </output>
  </function>


*/
