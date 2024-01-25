import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link,useLocation } from "react-router-dom";
import { Button,SIZE } from "baseui/button";
import { DatePicker } from "baseui/datepicker";
import { Table } from "baseui/table";




function DescribeWord(data,date){
  // console.dir(data)

  var dataResultArray = data ? (data.body.map(function(item) {
    return [item.word, item.count];
  })) : null;

  return (
    <>{data ? (
        <div>
          <Table 
            columns={["Word","Count"]} data={dataResultArray} 
          />
        </div>
      ) : (
        <div>
          <Table 
            columns={["Word","Count"]} data={[]} 
            isLoading
          />
        </div>
      )}
    </>
  )
}

export const Providers =() => {
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const providerid = query.get("id")
  const date_today = new Date().toISOString().slice(0,10)
  const [data, setData] = useState(null);
  const [date, setDate] = useState(date_today)

  console.log("hello")
  console.log(date)

  

  useEffect(() => {
    const fetchData = async () => {
      // const response_today = await axios.get("http://localhost:1323/frequency?Date="+date+"&Page=0&Limit=10");
      const response_today = await axios.get("https://murasa-nii.net/frequency?Date="+ date +"&Page=0&Limit=10");
      setData(response_today.data);
    }
    fetchData();
  }, [date]);

  const handleDateButtonClick = () => {
    const newDate = document.getElementById("date").value;
    setDate(newDate);
  };

  return (
    <>
      <h1>{providerid}</h1>
      <div>
        {DescribeWord(data,date)}
      </div>

      {/* todo */}
      <div>
      <DatePicker
      value={
        new Date(date)
      }
      onChange={({ date }) =>
      setDate(date.toISOString().slice(0,10) )
      }
      />
        <input type="text" id="date"></input>
        <Button
      onClick={handleDateButtonClick}
      size={SIZE.compact}
      >
        日付設定
      </Button>
        <input type="button" value="日付" onClick={handleDateButtonClick}></input>
      </div>

      <div>
        <a href="../">戻る</a>
      </div>
    </>
  );
}