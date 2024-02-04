import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link,useLocation } from "react-router-dom";
import { Button,SIZE } from "baseui/button";
import { DatePicker } from "baseui/datepicker";
import { Table } from "baseui/table";
import { Select } from "baseui/select";
import { Pagination } from "baseui/pagination";
import { Header } from "./Header.js";
import {
  TableBuilder,
  TableBuilderColumn,
} from 'baseui/table-semantic';


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


function DescribeProvider(providers){

  var dataResultArray = providers ? providers.body : null;

  console.log(dataResultArray)
  return (
    <>{dataResultArray ? (
      <div>
      <TableBuilder data={dataResultArray}>
        <TableBuilderColumn header="ID">
          {(row) => <a href={"./provider?id="+row.id}>{row.id}</a>}
        </TableBuilderColumn>
        <TableBuilderColumn header="SiteName" numeric>
          {(row) => row.site_name}
        </TableBuilderColumn>
      </TableBuilder>
    </div>
      ) : (
        <div>
        <Table 
          columns={["ID","SiteName"]} data={[]} 
          isLoading
        />
      </div>
      )}
    </>
  )
}

export const Providers = ({isDarkMode,setDarkFunc}) => {
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const providerid = query.get("id");
  const date_today = new Date().toISOString().slice(0,10);
  const [providers, setProviders] = useState(null);
  const [data, setData] = useState(null);
  const [date, setDate] = useState(date_today);
  const [limit, setLimit] = React.useState([
    {
      label: "10",
      id: "#F0F8FF"
    }
  ]);
  const [currentPage, setCurrentPage] = React.useState(1);

  console.log(date)

  

  useEffect(() => {
    console.log(limit)
    const fetchData = async () => {
      const lim = limit ? limit[0].label: "10";
      const page = currentPage ? currentPage - 1 : "0"
      const response = await axios.get("http://localhost:1323/provider?Date="+date+"&Page="+ page +"&Limit=" + lim + "&Provider=" + providerid);
      setProviders(response.data)
      // const response_today = await axios.get("https://murasa-nii.net/frequency?Date="+ date +"&Page="+ page +"&Limit=" + lim + "&Provider=" + providerid);
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
          <h1>Websites</h1>
      </header>
      <div style={{ display: "flex", justifyContent: "space-between", marginBottom: "10px" }}>
        <div style={{ width: "49%"}}>
          <span>表示数: </span>
          <div style={{  border: "2px solid black",
            borderRadius: "8px"}}>
          {SelectLimit(limit, setLimit)}
          </div>
        </div>
      </div>
      <div>
        {DescribeProvider(providers)}
      </div>

      <div>
        {GetPagination(currentPage,setCurrentPage)}
      </div>

      <div>
      <Button size={SIZE.compact} onClick={()=>{history.back()}}>
        戻る
      </Button>

      </div>
    </>
  );
}