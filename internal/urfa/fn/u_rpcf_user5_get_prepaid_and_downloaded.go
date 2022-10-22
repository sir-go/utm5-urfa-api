package fn

// rpcf_user5_get_prepaid_and_downloaded
func xU4027(c conn, _ Dict) Dict {
	c.Call(-0x4027)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4027" name="rpcf_user5_get_prepaid_and_downloaded">
  <input />
  <output>
    <integer name="bytes_in_mbyte" />
    <for count="1000" from="0" name="i">
      <integer name="unk1" />
      <if condition="eq" value="1" variable="unk1">
        <string name="iptraffic_service_name" />
        <for count="1000" from="0" name="j">
          <integer name="unk2" />
          <if condition="eq" value="1" variable="unk2">
            <integer array_index="i,j" name="old_tclass_id_array" />
            <string array_index="i,j" name="old_tclass_name_array" />
            <long array_index="i,j" name="old_prepaid_array" />
          </if>
          <if condition="eq" value="0" variable="unk2">
            <break />
          </if>
        </for>
        <for count="1000" from="0" name="j">
          <integer name="unk2" />
          <if condition="eq" value="1" variable="unk2">
            <integer array_index="i,j" name="tclass_id_array" />
            <string array_index="i,j" name="tclass_name_array" />
            <long array_index="i,j" name="prepaid_array" />
          </if>
          <if condition="eq" value="0" variable="unk2">
            <break />
          </if>
        </for>
        <for count="1000" from="0" name="j">
          <integer name="unk2" />
          <if condition="eq" value="1" variable="unk2">
            <integer array_index="i,j" name="downloaded_tclass_id_array" />
            <string array_index="i,j" name="downloaded_tclass_name_array" />
            <long array_index="i,j" name="downloaded_array" />
          </if>
          <if condition="eq" value="0" variable="unk2">
            <break />
        </if>
        </for>
      </if>
      <if condition="eq" value="0" variable="unk1">
        <break />
      </if>
    </for>
  </output>
  </function>


*/
