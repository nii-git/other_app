import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { SamplePage1 } from "./components/SamplePage1.js";
import { SamplePage2 } from "./components/SamplePage2.js";
import { NotFound } from "./components/NotFound.js";
import  TopPage  from "./components/TopPage.js";

export const RouterConfig =() => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<TopPage />} />
        <Route path="page1" element={<SamplePage1 />} />
        <Route path="page2" element={<SamplePage2 />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}