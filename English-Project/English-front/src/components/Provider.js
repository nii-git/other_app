import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link,useLocation } from "react-router-dom";
import { Button,SIZE } from "baseui/button";
import { DatePicker } from "baseui/datepicker";
import { Table } from "baseui/table";
import { Select } from "baseui/select";
import { Pagination } from "baseui/pagination";
import { Header } from "./Header.js";


function SelectLimit(limit,setFunc)  {

  return (
    <Select
      clearable={false}
      options={[
        {
          label: "10",
          id: "#F0F8FF"
        },
        {
          label: "100",
          id: "#FAEBD7"
        },
        {
          label: "1000",
          id: "#00FFFF"
        }
      ]}
      value={limit}
      placeholder="Select color"
      onChange={params => setFunc(params.value)}
    />
  );
}


function GetPagination (page,setFunc) {
  return (
    <Pagination
      numPages={20}
      size={SIZE.compact}
      currentPage={page}
      onPageChange={({ nextPage }) => {
        setFunc(
          Math.min(Math.max(nextPage, 1), 20)
        );
      }}
    />
  );
}


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

export const Provider = ({isDarkMode,setDarkFunc}) => {
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const providerid = query.get("id");
  const date_today = new Date().toISOString().slice(0,10);
  const [data, setData] = useState(null);
  const [date, setDate] = useState(date_today);
  const [limit, setLimit] = React.useState([
    {
      label: "10",
      id: "#F0F8FF"
    }
  ]);
  const [currentPage, setCurrentPage] = React.useState(1);

  console.log("hello")
  console.log(date)

  

  useEffect(() => {
    console.log(limit)
    const fetchData = async () => {
      const lim = limit ? limit[0].label: "10";
      const page = currentPage ? currentPage - 1 : "0"
      // const response_today = await axios.get("http://localhost:1323/frequency?Date="+date+"&Page=0&Limit=10");
      setData(null)
      const response_today = await axios.get("https://murasa-nii.net/frequency?Date="+ date +"&Page="+ page +"&Limit=" + lim + "&Provider=" + providerid);
      setData(response_today.data);
    }
    fetchData();
  }, [date,limit,currentPage]);

  const handleDateButtonClick = () => {
    const newDate = document.getElementById("date").value;
    setDate(newDate);
  };

  return (
    <>
      <Header></Header>

      <header className='about-header'>
          <h1>{providerid}</h1>
      </header>

      <div>
      <DatePicker
      value={
        new Date(date)
      }
      onChange={({ date }) =>
      setDate(date.toISOString().slice(0,10) )
      }
      />
      </div>

      <div>
        {SelectLimit(limit,setLimit)}
      </div>

      <div>
        {GetPagination(currentPage,setCurrentPage)}
      </div>

      <div>
        {DescribeWord(data,date)}
      </div>

      <div>
        {GetPagination(currentPage,setCurrentPage)}
      </div>

      <div>
      <Button size={SIZE.compact} onClick={()=>{history.back()}}>
        戻る
      </Button>
      <Button size={SIZE.compact} onClick={()=>{setDarkFunc(!isDarkMode)}}>
        change
      </Button>
      </div>
    </>
  );
}