import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link,useLocation } from "react-router-dom";




function DescribeWord(data,date){
  console.dir(data)

  return (
    <>{data ? (
        <div>
          <table>
            <thead>
              <tr>
                <th colSpan="2">{date} 集計結果</th>
              </tr>
            </thead>
            <tbody>
              {data.body.map((item, index) => (
                <tr>
                  <td key={"word"+index}>{item.word}</td>
                  <td key={"count"+index}>{item.count}</td>
                </tr>
            ))}
            </tbody>
          </table>
        </div>
      ) : (
        <p>Today's Loading data...</p>
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
  

  useEffect(() => {
    const fetchData = async () => {
      const response_today = await axios.get("http://localhost:1323/frequency?Date="+date+"&Page=0&Limit=10");
      // const response_today = await axios.get("https://murasa-nii.net/frequency?Date="+ date_today+"&Page=0&Limit=10");
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
      <h1>Sample Home</h1>
      <>
        {providerid}
      </>
      <div>
        {data ? DescribeWord(data,date) : "データ取得中です..."}
      </div>

      {/* todo */}
      <div>
        <input type="text" id="date"></input>
        <input type="button" value="日付" onClick={handleDateButtonClick}></input>
      </div>
    </>
  );
}