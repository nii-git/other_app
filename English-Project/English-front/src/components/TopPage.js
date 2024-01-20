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
        <p>Today's Loading data...</p>
      )}
    </>
  )
}

function DescribeProvider(data){
  console.log("DescribeProvider")
  console.dir(data)

  return (
    <>{data && data.body  ? (
        <div>
          <table>
            <thead>
              <tr>
                <th colspan="2">providers</th>
              </tr>
            </thead>
            <tbody>
              {data.body.map((item, index) => (
                <tr>
                  <a href={"./providers?id=" +item.id }>{item.site_name}</a>
                  <td key={"word"+index}>{item.id}</td>
                  <td key={"count"+index}>{item.site_name}</td>
                  <td key={"url"+index}>{item.url}</td>
                </tr>
            ))}
            </tbody>
          </table>
        </div>
      ) : (
        <p>Provider Loading data...</p>
      )}
    </>
  )
}

const TopPage = () => {
  const [data_today, setData_today] = useState(null);
  const [provider, setProvider] = useState(null);
  const date_today = new Date().toISOString().slice(0,10)

  useEffect(() => {
    const fetchData = async () => {
      
      try {
        const response_provider = await axios.get("http://localhost:1323/provider")
        setProvider(response_provider.data);
        const response_today = await axios.get("http://localhost:1323/frequency?Date="+"2024-01-15"+"&Page=0&Limit=10");
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
      <div>
        {data_today ? DescribeToday(data_today,date_today) : "データ取得中です..."}
      </div>
      <div>
        {provider ? DescribeProvider(provider): "Providerデータ取得中です..."}
      </div>
    </div>
  );
};

export default TopPage;
