// TopPage.js

import React, { useEffect, useState } from "react";
import axios from "axios";

function DescribeToday(data,date){
  console.dir(data)

  return (
    <>{data ? (
        <div>
          <table>
            <thead>
              <tr>
                <th colspan="2">{date} 集計結果</th>
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
        <p>Loading data...</p>
      )}
    </>
  )
}

const TopPage = () => {
  const [data_today, setData_today] = useState(null);
  const date_today = new Date().toISOString().slice(0,10)

  useEffect(() => {
    const fetchData = async () => {
      
      try {
        const response_today = await axios.get("http://localhost:1323/frequency?Date="+ date_today+"&Page=0&Limit=10");
        // const response_today = await axios.get("https://murasa-nii.net/frequency?Date="+ date_today+"&Page=0&Limit=10");
        setData_today(response_today.data);
      } catch (error) {
        console.error("Error fetching data:", error);
        const errMsg = error.response.data.error || error.messaage
        setData_today(errMsg)
        console.error("Error fetching data:", errMsg);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      {data_today ? DescribeToday(data_today,date_today) : "データ取得中です..."}
    </div>
  );
};

export default TopPage;
