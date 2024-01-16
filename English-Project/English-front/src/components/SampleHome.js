import React from "react";
import { Link } from "react-router-dom";

export const SampleHome =() => {
  return (
    <>
      <h1>Sample Home</h1>
      <nav>
        <ul>
          <li><Link to="page1">Sample Page1</Link></li>
          <li><Link to="page2">Sample Page2</Link></li>
        </ul>
      </nav>  
    </>
  );
}