import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { SamplePage1 } from "./components/SamplePage1.js";
import { SamplePage2 } from "./components/SamplePage2.js";
import { NotFound } from "./components/NotFound.js";
import { Provider } from "./components/Provider.js";
import  TopPage   from "./components/TopPage.js";
import  { About }  from "./components/About.js";
import {Providers} from "./components/Providers.js"


export const RouterConfig = ({isDarkMode,setDarkFunc}) => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<TopPage />} />
        {/* darkmodeいらないかも */}
        <Route path="provider" element={<Provider isDarkMode={isDarkMode} setDarkFunc={setDarkFunc} />} />
        <Route path="providers" element={<Providers isDarkMode={isDarkMode} setDarkFunc={setDarkFunc} />} />
        <Route path="about" element={<About />} />
        <Route path="page1" element={<SamplePage1 />} />
        <Route path="page2" element={<SamplePage2 />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}